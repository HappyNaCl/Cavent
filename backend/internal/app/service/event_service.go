package service

import (
	"context"
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/HappyNaCl/Cavent/backend/internal/app/mapper"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/cache"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	rediscache "github.com/HappyNaCl/Cavent/backend/internal/infrastructure/cache/redis"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb/paginate"
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
	categoryRepo repo.CategoryRepository
	favoriteRepo repo.FavoriteRepository

	userInterestCache cache.UserInterestCache
	categoriesCache cache.CategoriesCache

	asynqClient *asynq.Client
}

func NewEventService(db *gorm.DB, redis *redis.Client, client *asynq.Client) *EventService {
	return &EventService{
		eventRepo: postgresdb.NewEventGormRepo(db),
		userRepo: postgresdb.NewUserGormRepo(db),
		categoryRepo: postgresdb.NewCategoryGormRepo(db),
		favoriteRepo: postgresdb.NewFavoriteGormRepo(db),

		userInterestCache: rediscache.NewUserInterestCache(redis),
		categoriesCache: rediscache.NewCategoriesCache(redis),
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

	eventResult := mapper.NewEventResultFromEvent(eventModel, false)
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
	if com.UserId == nil {
		for _, event := range events {
			eventResults = append(eventResults, mapper.NewBriefEventResultFromEvent(event, false))
		}
	}else {
		eventIds := e.GetEventIds(events)
		isFavorites, err := e.checkFavorited(*com.UserId, eventIds)
		if err != nil {
			return nil, err
		}

		for _, event := range events {
			var favorited bool
			if isFavorite, exists := isFavorites[event.Id]; exists {
				favorited = isFavorite
			} else {
				favorited = false
			}
			zap.L().Sugar().Infof("%v", isFavorites)
			eventResults = append(eventResults, mapper.NewBriefEventResultFromEvent(event, favorited))
		}
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

	categoryIds := make([]uuid.UUID, 0, len(categories))
	for _, category := range categories {
		categoryIds = append(categoryIds, category.Id)
	}

	events, err := e.eventRepo.GetEventsByCategories(categoryIds, 8)
	if err != nil {
		return nil, err
	}

	eventIds := e.GetEventIds(events)
	isFavorites, err := e.checkFavorited(com.UserId, eventIds)
	if err != nil {
		return nil, err
	}

	eventResults := make([]*common.BriefEventResult, 0, len(events))
	for _, event := range events {
		var favorited bool
		if isFavorite, exists := isFavorites[event.Id]; exists {
			favorited = isFavorite
		} else {
			favorited = false
		}

		eventResults = append(eventResults, mapper.NewBriefEventResultFromEvent(event, favorited))
	}

	return &command.GetUserInterestedEventsCommandResult{
		Result: eventResults,
	}, nil
}

func (e EventService) GetRandomEvents(ctx context.Context, com *command.GetRandomEventsCommand) (*command.GetRandomEventsCommandResult, error) {
	categories, err := e.categoriesCache.GetAllCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("cache lookup failed: %w", err)
	}
	if categories == nil {
		zap.L().Sugar().Info("Cache miss for all categories")
		
		categoriesResult, err := e.categoryRepo.GetAllCategory()
		if err != nil {
			return nil, fmt.Errorf("db lookup failed: %w", err)
		}
		categories = mapper.NewCategoryResultsFromCategoryType(categoriesResult)
		err = e.categoriesCache.SetAllCategories(ctx, categories)
		if err != nil {
			return nil, fmt.Errorf("failed to set categories in cache: %w", err)
		}
		zap.L().Sugar().Infof("Cache miss for categories:all, cached %d categories", len(categories))
	}else {
		zap.L().Sugar().Info("Cache hit for categories:all")
	}

	randomCategoryAmount := 3

	categoryIds := make([]uuid.UUID, 0, randomCategoryAmount)
	for range randomCategoryAmount {
		categoryIds = append(categoryIds, categories[rand.IntN(len(categories))].Id)
		categories = append(categories[:len(categories)-1], categories[len(categories):]...)
	}

	events, err := e.eventRepo.GetEventsByCategories(categoryIds, 8)
	if err != nil {
		return nil, err
	}

	eventResults := make([]*common.BriefEventResult, 0, len(events))
	for _, event := range events {
		eventResults = append(eventResults, mapper.NewBriefEventResultFromEvent(event, false))
	}

	return &command.GetRandomEventsCommandResult{
		Result: eventResults,
	}, nil
}

func (e EventService) GetEventDetails(com *command.GetEventDetailsCommand) (*command.GetEventDetailsCommandResult, error) {
	event, err := e.eventRepo.GetEventByID(com.EventID)
	if err != nil {
		return nil, err
	}

	if event == nil {
		return nil, fmt.Errorf("event not found with ID: %s", com.EventID)
	}

	var favorited bool = false
	if com.UserId != nil {
		isFavorites, err := e.checkFavorited(*com.UserId, []uuid.UUID{event.Id})
		if err != nil {
			return nil, err
		}
		favorited = isFavorites[event.Id]
	}

	eventResult := mapper.NewEventResultFromEvent(event, favorited)


	asynqTask, err := tasks.NewEventViewPayload(com.UserId, com.EventID)
	if err != nil {
		return nil, err
	}
	
	_, err = e.asynqClient.Enqueue(asynqTask, asynq.MaxRetry(5), asynq.Queue(tasks.TypeEventView))
	if err != nil {
		zap.L().Sugar().Errorf("Failed to enqueue event view task: %v", err)
		return nil, err
	}

	return &command.GetEventDetailsCommandResult{
		Result: eventResult,
	}, nil
}

func (e EventService) GetEventsByCampus(com *command.GetCampusEventsCommand) (*command.GetCampusEventsCommandResult, error) {
	results, err := e.eventRepo.GetCampusEvents(com.CampusId, paginate.Pagination{
		Limit: com.Limit,
		Page:  com.Page,
	})
	if err != nil {
		return nil, err
	}

	events := make([]*common.BriefEventResult, 0, len(results.Rows.([]*model.Event)))
	for _, event := range results.Rows.([]*model.Event) {
		var favorited bool = false
		if com.UserId != nil {
			isFavorites, err := e.checkFavorited(*com.UserId, []uuid.UUID{event.Id})
			if err != nil {
				return nil, err
			}
			favorited = isFavorites[event.Id]
		}

		eventResult := mapper.NewBriefEventResultFromEvent(event, favorited)
		events = append(events, eventResult)
	}

	return &command.GetCampusEventsCommandResult{
		Result:     events,
		TotalRows:  results.TotalRows,
		TotalPages: results.TotalPages,
		Page:       results.Page,
		Limit:      results.Limit,
	}, nil
}


func (e EventService) checkFavorited(userId string, eventIds []uuid.UUID) (map[uuid.UUID]bool, error) {
	results, err := e.favoriteRepo.IsFavorited(userId, eventIds)
	if err != nil {
		zap.L().Sugar().Errorf("Failed to check favorited status: %v", err)
		return nil, err
	}

	return results, nil
}

func (e EventService) GetEventIds(events []*model.Event) []uuid.UUID {
	eventIds := make([]uuid.UUID, 0, len(events))
	for _, event := range events {
		eventIds = append(eventIds, event.Id)
	}
	return eventIds
}

func (e EventService) SearchEvents(com *command.GetSearchEventCommand) (*command.GetSearchEventCommandResult, error) {
	events, err := e.eventRepo.SearchEvents(com.Query)
	if err != nil {
		return nil, err
	}

	eventResults := make([]*common.EventSearchResult, 0, len(events))
	for _, event := range events {
		eventResults = append(eventResults, mapper.NewSearchResultFromEvent(event))
	}

	return &command.GetSearchEventCommandResult{
		Result: eventResults,
	}, nil
}

func (e EventService) GetAllEvents(com *command.GetAllEventCommand) (*command.GetAllEventCommandResult, error) {
	page := paginate.Pagination{
		Limit: com.Limit,
		Page: com.Page,
		Sort: "start_time ASC",
		Filter: com.Filters,
		FilterArgs: com.FilterArgs,
	}

	results, err := e.eventRepo.GetAllEvents(page)
	if err != nil {
		return nil, err
	}
	
	events := make([]*common.BriefEventResult, 0, len(results.Rows.([]*model.Event)))
	for _, event := range results.Rows.([]*model.Event) {
		var favorited bool = false
		if com.UserId != nil {
			isFavorites, err := e.checkFavorited(*com.UserId, []uuid.UUID{event.Id})
			if err != nil {
				return nil, err
			}
			favorited = isFavorites[event.Id]
		}

		eventResult := mapper.NewBriefEventResultFromEvent(event, favorited)
		events = append(events, eventResult)
	}

	return &command.GetAllEventCommandResult{
		Limit: results.Limit,
		Page: results.Page,
		TotalPage: results.TotalPages,
		TotalRows: int(results.TotalRows),
		Result: events,
	}, nil
}

func (e EventService) GetUserFavoritedEvent (com *command.GetUserFavoritedEventCommand) (*command.GetUserFavoritedEventResult, error){
	events, err := e.eventRepo.GetUserFavoritedEvent(com.UserId)
	if err != nil {
		return nil, err
	}
	eventResults := make([]*common.BriefEventResult, 0, len(events))
	for _, event := range events {
		eventResults = append(eventResults, mapper.NewBriefEventResultFromEvent(event, true))
	}
	return &command.GetUserFavoritedEventResult{
		Result: eventResults,
	}, nil
}