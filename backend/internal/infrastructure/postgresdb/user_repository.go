package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserGormRepo struct {
	db    *gorm.DB
	redis *redis.Client
	logger *zap.Logger
}

// FindByEmail implements repo.UserRepository.
func (u *UserGormRepo) FindByEmail(email string) (*model.User, error) {
	panic("unimplemented")
}

func NewUserGormRepo(db *gorm.DB, redis *redis.Client, logger *zap.Logger) repo.UserRepository {
	return &UserGormRepo{
		db:    db,
		redis: redis,
		logger: logger,
	}
}
