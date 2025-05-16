package app

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
	redis *redis.Client
}

func NewUserService(db *gorm.DB, redis *redis.Client) *UserService {
	return &UserService{
		db: db,
		redis: redis,
	}
}

