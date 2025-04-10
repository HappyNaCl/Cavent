package application

import (
	"github.com/HappyNaCl/Cavent/src/config"
	"github.com/HappyNaCl/Cavent/src/domain"
	"github.com/HappyNaCl/Cavent/src/domain/factory"
	"github.com/HappyNaCl/Cavent/src/infrastructure/persistence"
)

func RegisterUser(fullName string, email string, password string) (*domain.User, error) {
	user := factory.UserFactory().GetUser(fullName, email, password)
	
	return persistence.UserRepository(config.Database).RegisterUser(user)
}