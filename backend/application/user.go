package application

import (
	"errors"
	str "strings"

	"github.com/HappyNaCl/Cavent/backend/config"
	"github.com/HappyNaCl/Cavent/backend/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/domain/model"
	"github.com/HappyNaCl/Cavent/backend/infrastructure/persistence"
)

func RegisterUser(fullName string, email string, password string) (*model.User, error) {
	if fullName == "" || email == "" || password == "" {
		return nil, errors.New("full name, email and password are required")
	}

	if !str.Contains(email, "@") || !str.Contains(email, ".") {
		return nil, errors.New("invalid email format")
	}

	if len(password) < 8 {
		return nil, errors.New("password must be at least 8 characters long")
	}

	user := factory.UserFactory().GetUser(fullName, email, password)
	
	return persistence.UserRepository(config.Database).RegisterUser(user)
}

func UpdatePrefences(userId string, preferences []string) error {
	if userId == "" {
		return errors.New("user ID is required")
	}

	if len(preferences) == 0 {
		return errors.New("preferences are required")
	}

	return persistence.UserRepository(config.Database).UpdateInterest(userId, preferences)
}

func GetUserTag(userId string) ([]model.Tag, error) {
	if userId == "" {
		return nil, errors.New("user ID is required")
	}

	tags, err := persistence.UserRepository(config.Database).GetUserTag(userId)
	if err != nil {
		return nil, err
	}

	return tags, nil
}