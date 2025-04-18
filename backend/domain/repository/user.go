package repository

import (
	"github.com/HappyNaCl/Cavent/backend/domain/model"
	"github.com/markbates/goth"
)

type UserRepository interface {
	FindByProviderID(providerID string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	RegisterUser(user *model.User) (*model.User, error)
	RegisterOrLoginOauthUser(gothUser goth.User, provider string) (*model.User, error)
	UpdateInterest(userId string, interestIds []string) error
	GetUserTag(userId string) ([]model.Tag, error)
}