package app

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/postgresdb"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct {
	logger *zap.Logger
	userRepo repo.UserRepository
}

func NewUserService(db *gorm.DB, redis *redis.Client, logger *zap.Logger) *UserService {
	return &UserService{
		logger: logger,
		userRepo: postgresdb.NewUserGormRepo(db, redis, logger),
	}
}

