package app

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)


type AuthService struct {
	
}

func NewAuthService(db *gorm.DB, redis *redis.Client) *AuthService {
	return &AuthService{

	}
}
