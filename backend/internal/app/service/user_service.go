package service

import (
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

