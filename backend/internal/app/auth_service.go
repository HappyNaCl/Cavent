package app

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/postgresdb"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)


type AuthService struct {
	authRepo repo.AuthRepository
	logger *zap.Logger
}

func NewAuthService(db *gorm.DB, redis *redis.Client, logger *zap.Logger) *AuthService {
	return &AuthService{
		authRepo: postgresdb.NewAuthGormRepo(db, redis, logger),
		logger: logger,
	}
}
