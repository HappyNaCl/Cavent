package repository

import (
	"github.com/HappyNaCl/Cavent/src/domain"
	"github.com/markbates/goth"
)

// Make a repo interface containing methods to query the database for user
type UserRepository interface {
	FindByProviderID(providerID string) (*domain.User, error)
	RegisterUser(user *domain.User) error
	RegisterOrLoginUser(gothUser goth.User, provider string) (*domain.User, error)
}