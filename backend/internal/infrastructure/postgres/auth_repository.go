package postgres

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthGormRepo struct {
	db *gorm.DB
	redis *redis.Client
}

func NewAuthGormRepo(db *gorm.DB, redis *redis.Client) repo.AuthRepository {
	return &AuthGormRepo{
		db: db,
		redis: redis,
	}
}

