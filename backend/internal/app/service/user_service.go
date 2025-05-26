package service

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/mapper"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/errors"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo repo.UserRepository
}

func NewUserService(db *gorm.DB, redis *redis.Client) *UserService {
	return &UserService{
		userRepo: postgresdb.NewUserGormRepo(db),
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

func (us *UserService) GetUserInterests(com *command.GetUserInterestsCommand) (*command.GetUserInterestsCommandResult, error) {
	interests, err := us.userRepo.GetUserInterests(com.UserId)
	if err != nil {
		return nil, err
	}

	return &command.GetUserInterestsCommandResult{
	  CategoryTypes	: mapper.NewCategoriesResultFromCategory(interests),
	}, nil
}

func (us *UserService) UpdateUserInterests(com *command.UpdateUserInterestCommand) (*command.UpdateUserInterestCommandResult, error) {
	if len(com.CategoryIds) <= 0 {
		return nil, errors.ErrInterestLength
	}

	err := us.userRepo.UpdateUserInterests(com.UserId, com.CategoryIds)
	if err != nil {
		return nil, err
	}

	return &command.UpdateUserInterestCommandResult{}, nil
}