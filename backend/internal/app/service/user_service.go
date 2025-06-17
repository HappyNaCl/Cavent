package service

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/mapper"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/cache"
	domainerrors "github.com/HappyNaCl/Cavent/backend/internal/domain/errors"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	rediscache "github.com/HappyNaCl/Cavent/backend/internal/infrastructure/cache/redis"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/queue/tasks"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo repo.UserRepository
	campusRepo repo.CampusRepository

	asynqClient *asynq.Client

	userInterestCache cache.UserInterestCache
}

func NewUserService(db *gorm.DB, redis *redis.Client, asynq *asynq.Client) *UserService {
	return &UserService{
		userRepo: postgresdb.NewUserGormRepo(db),
		campusRepo: postgresdb.NewCampusGormRepo(db),
		asynqClient: asynq,
		userInterestCache: rediscache.NewUserInterestCache(redis),
	}
}

func (us *UserService) GetBriefUser(com *command.GetBriefUserCommand) (*command.GetBriefUserCommandResult, error) {
	user, err := us.userRepo.GetBriefUser(com.UserId)
	if err != nil {
		return nil, err
	}

	userResult := mapper.NewUserResultFromUser(user).ToBrief()

	return &command.GetBriefUserCommandResult{
		Result: userResult,
	}, nil
}

func (us *UserService) GetUserInterests(ctx context.Context, com *command.GetUserInterestsCommand) (*command.GetUserInterestsCommandResult, error) {
	categories, err := us.userInterestCache.GetUserInterest(ctx, com.UserId)
	if err != nil {
		return nil, fmt.Errorf("cache lookup failed: %w", err)
	}
	if categories != nil {
		zap.L().Sugar().Infof("Cache hit for user interests: %s", com.UserId)
		return &command.GetUserInterestsCommandResult{
			Result: categories,
		}, nil
	}

	interests, err := us.userRepo.GetUserInterests(com.UserId)
	if err != nil {
		return nil, fmt.Errorf("db lookup failed: %w", err)
	}

	result := mapper.NewCategoriesResultFromCategory(interests)

	err = us.userInterestCache.SetUserInterest(ctx, com.UserId, result)
	if err != nil {
		return nil, fmt.Errorf("failed to set user interests in cache: %w", err)
	}

	zap.L().Sugar().Infof("Cache miss for user interests: %s, cached %d categories", com.UserId, len(result))
	return &command.GetUserInterestsCommandResult{
		Result: result,
	}, nil
}


func (us *UserService) UpdateUserInterests(com *command.UpdateUserInterestCommand) (*command.UpdateUserInterestCommandResult, error) {
	if len(com.CategoryIds) <= 0 {
		return nil, domainerrors.ErrInterestLength
	}

	err := us.userRepo.UpdateUserInterests(com.UserId, com.CategoryIds)
	if err != nil {
		return nil, err
	}	

	err = us.userInterestCache.Invalidate(context.Background(), com.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to invalidate user interests cache: %w", err)
	}
	return &command.UpdateUserInterestCommandResult{}, nil
}

func (us *UserService) UpdateUserCampus(com *command.UpdateUserCampusCommand) (*command.UpdateUserCampusCommandResult, error) {
	campus, err := us.campusRepo.FindCampusByInviteCode(com.InviteCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrCampusNotFound
		}
		return nil, err
	}

	zap.L().Sugar().Infof("Updating user campus for user %s to campus %s", com.UserId, campus.Id)
	user, err := us.userRepo.UpdateUserCampus(com.UserId, campus.Id)
	if err != nil {
		return nil, err
	}

	userResult := mapper.NewUserResultFromUser(user).ToBrief()
	return &command.UpdateUserCampusCommandResult{
		User: &userResult,
	}, nil
}

func (us *UserService) GetUserProfile(ctx context.Context, com *command.GetUserProfileCommand) (*command.GetUserProfileCommandResult, error) {
	user, err := us.userRepo.GetUserProfile(com.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrUserNotFound
		}
		return nil, err
	}

	userResult := mapper.NewUserResultFromUser(user).ToProfile()
	return &command.GetUserProfileCommandResult{
		Result: userResult,
	}, nil
}

func (us *UserService) UpdateUserProfile(ctx context.Context, com *command.UpdateUserProfileCommand) (*command.UpdateUserProfileCommandResult, error) {	
	factory := factory.NewUserProfileFactory()
	user:= factory.CreateUserProfileResult(
		com.UserId,
		com.Name,
		com.Description,
		com.PhoneNumber,
		com.Address,
	)

	if com.AvatarBytes != nil && com.AvatarExt != nil {
		publicUrl := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", 
								os.Getenv("SUPABASE_PROJECT_URL"), 
								os.Getenv("SUPABASE_BUCKET_NAME"), 
								"profile/" + com.UserId + *com.AvatarExt)
		user.AvatarUrl = publicUrl
	}

	updatedUser, err := us.userRepo.UpdateUserProfile(user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrUserNotFound
		}
		return nil, err
	}

	if com.AvatarBytes != nil && com.AvatarExt != nil {
		asynqTask, err := tasks.NewImageUploadTask(com.AvatarBytes, *com.AvatarExt, "profile/" + com.UserId + *com.AvatarExt)
		if err != nil {
			return nil, err
		}

		_, err = us.asynqClient.Enqueue(asynqTask, asynq.MaxRetry(5), asynq.Queue(tasks.TypeImageUpload), )
		if err != nil {
			return nil, err
		}
	}


	return &command.UpdateUserProfileCommandResult{
		Result: mapper.NewUserResultFromUser(updatedUser).ToProfile(),
	}, nil
}