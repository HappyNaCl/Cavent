package postgresdb

import (
	"errors"
	"strings"

	"github.com/HappyNaCl/Cavent/backend/internal/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/markbates/goth"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthGormRepo struct {
	db    *gorm.DB
}

var (
	ErrDuplicateEmail = errors.New("email already exists")
)

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
	var user model.User

	result := a.db.Where("id = ? AND provider = ?", gothUser.UserID, provider).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			userModel := factory.NewUserFactory().GetOAuthUser(provider, gothUser.UserID, strings.Split(gothUser.Email, "@")[0], gothUser.Email, "", gothUser.AvatarURL)
			
			zap.L().Sugar().Infof("[INFO] User not found in database, creating new user")
			
			err := a.db.Create(&userModel).Error
			if err != nil {
				return nil, err
			}

			return userModel, nil
		} else {
			return nil, result.Error
		}
	}

	zap.L().Sugar().Infof("[INFO] User found in database")
	return &user, nil
}

// RegisterUser implements repo.AuthRepository.
func (a *AuthGormRepo) RegisterUser(user *model.User) (*model.User, error) {
	err := a.db.Create(user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewAuthGormRepo(db *gorm.DB) repo.AuthRepository {
	return &AuthGormRepo{
		db:    db,
	}
}
