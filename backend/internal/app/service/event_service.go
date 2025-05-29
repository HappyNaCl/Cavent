package service

import (
	"context"
	"fmt"
	"os"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/HappyNaCl/Cavent/backend/internal/app/mapper"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/cache"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	rediscache "github.com/HappyNaCl/Cavent/backend/internal/infrastructure/cache/redis"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/queue/tasks"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type EventService struct {
	eventRepo repo.EventRepository
	userRepo repo.UserRepository
	userInterestCache cache.UserInterestCache
	asynqClient *asynq.Client
}

func NewEventService(db *gorm.DB, redis *redis.Client, client *asynq.Client) *EventService {
	return &EventService{
		eventRepo: postgresdb.NewEventGormRepo(db),
		userRepo: postgresdb.NewUserGormRepo(db),
		userInterestCache: rediscache.NewUserInterestCache(redis),
		asynqClient: client,
	}
}


func (e EventService) CreateEvent(com *command.CreateEventCommand) (*command.CreateEventCommandResult, error) {
	eventId := uuid.New()

	publicUrl := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", 
							 os.Getenv("SUPABASE_PROJECT_URL"), 
							 os.Getenv("SUPABASE_BUCKET_NAME"), 
							 "events/" + eventId.String() + com.BannerExt)

	factory := factory.NewEventFactory()
	event := factory.GetEvent(
		eventId,
		com.CategoryId,
		com.CreatedById,
		com.Title,
		com.EventType,
		com.TicketType,
		com.Location,
		publicUrl,
		com.StartTime.Unix(),
		com.Ticket,
	)

	if com.Description != nil {
		event.Description = com.Description
	}

	if com.EndTime != nil {
		event.EndTime = com.EndTime
	} 

	campusId, err := e.userRepo.GetCampusId(com.CreatedById)
	if err != nil {
		return nil, err
	}
	event.CampusId = *campusId
	zap.L().Sugar().Infof("%s", event.CategoryId)
	eventModel, err := e.eventRepo.CreateEvent(event)
	if err != nil {
		return nil, err
	}

	asynqTask, err := tasks.NewImageUploadTask(com.BannerBytes, com.BannerExt, "events/" + eventId.String() + com.BannerExt)
	if err != nil {
		return nil, err
	}

	_, err = e.asynqClient.Enqueue(asynqTask, asynq.MaxRetry(5), asynq.Queue(tasks.TypeImageUpload), )
	if err != nil {
		return nil, err
	}

	eventResult := mapper.NewEventResultFromEvent(eventModel)
	return &command.CreateEventCommandResult{
		Result: eventResult,
	}, nil
}

func (e EventService) GetEvents(com *command.GetEventsCommand) (*command.GetEventsCommandResult, error) {
	events, err := e.eventRepo.GetEvents(com.Limit)
	if err != nil {
		return nil, err
	}

	eventResults := make([]*common.BriefEventResult, 0, len(events))
	for _, event := range events {
		eventResults = append(eventResults, mapper.NewBriefEventResultFromEvent(event))
	}

	return &command.GetEventsCommandResult{
		Result: eventResults,
	}, nil
}

func (e EventService) GetUserInterestedEvents(ctx context.Context, com *command.GetUserInterestedEventsCommand) (*command.GetUserInterestedEventsCommandResult, error) {
	categories, err := e.userInterestCache.GetUserInterest(ctx, com.UserId)
	if err != nil {
		return nil, fmt.Errorf("cache lookup failed: %w", err)
	}

	if categories == nil {
		zap.L().Sugar().Infof("Cache miss for user interests: %s", com.UserId)

		interests, err := e.userRepo.GetUserInterests(com.UserId)
		if err != nil {
			return nil, fmt.Errorf("db lookup failed: %w", err)
		}

		categories = mapper.NewCategoriesResultFromCategory(interests)

		_ = e.userInterestCache.SetUserInterest(ctx, com.UserId, categories)
	} else {
		zap.L().Sugar().Infof("Cache hit for user interests: %s", com.UserId)
	}

	// Proceed to fetch events based on categories
	categoryIds := make([]uuid.UUID, 0, len(categories))
	for _, category := range categories {
		categoryIds = append(categoryIds, category.Id)
	}

	events, err := e.eventRepo.GetEventsByCategories(categoryIds, 8)
	if err != nil {
		return nil, err
	}

	eventResults := make([]*common.BriefEventResult, 0, len(events))
	for _, event := range events {
		eventResults = append(eventResults, mapper.NewBriefEventResultFromEvent(event))
	}

	return &command.GetUserInterestedEventsCommandResult{
		Result: eventResults,
	}, nil
}