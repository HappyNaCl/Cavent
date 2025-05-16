package repo

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/markbates/goth"
)

type AuthRepository interface {
	LoginUser(email string) (*model.User, error)
	RegisterUser(user *model.User) (*model.User, error)
	RegisterOrLoginOauthUser(gothUser goth.User, provider string) (*model.User, error)
}