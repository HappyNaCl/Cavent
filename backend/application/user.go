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