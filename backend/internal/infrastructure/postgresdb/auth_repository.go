package postgresdb

import (
	"log"
	"strings"

	"github.com/HappyNaCl/Cavent/backend/internal/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/markbates/goth"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthGormRepo struct {
	db    *gorm.DB
	redis *redis.Client
	logger *zap.Logger
}

// LoginUser implements repo.AuthRepository.
func (a *AuthGormRepo) LoginUser(email string) (*model.User, error) {
	var user model.User
	err := a.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// RegisterOrLoginOauthUser implements repo.AuthRepository.
func (a *AuthGormRepo) RegisterOrLoginOauthUser(gothUser goth.User, provider string) (*model.User, error) {
	var user *model.User

	result := a.db.Where("id = ? AND provider = ?", gothUser.UserID, provider).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			user = factory.NewUserFactory().GetOAuthUser(provider, gothUser.UserID, strings.Split(gothUser.Email, "@")[0], gothUser.Email, "", gothUser.AvatarURL)
			log.Println("[INFO] User not found in database, creating new user:", user)
			a.db.Create(&user)
		} else {
			return nil, result.Error
		}
	}

	log.Println("[INFO] User found in database:", user)
	return user, nil
}

// RegisterUser implements repo.AuthRepository.
func (a *AuthGormRepo) RegisterUser(user *model.User) (*model.User, error) {
	err := a.db.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewAuthGormRepo(db *gorm.DB, redis *redis.Client, logger *zap.Logger) repo.AuthRepository {
	return &AuthGormRepo{
		db:    db,
		redis: redis,
		logger: logger,
	}
}
