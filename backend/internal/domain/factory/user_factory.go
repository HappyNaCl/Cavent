package factory

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserFactory struct {}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func (u *UserFactory) GetOAuthUser(provider string, providerId string, name string, email string, password string, avatarUrl string) *model.User {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil
	}

	return &model.User{
		Id: uuid.New(),
		Provider: provider,
		ProviderId: providerId,
		Name: name,
		Email: email,
		Password: string(hash),
		AvatarUrl: avatarUrl,
	}
}

func (u *UserFactory) GetUser(name string, email string, password string) *model.User {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil
	}

	return &model.User{
		Id: uuid.New(),
		Provider: "credential",
		Name: name,
		Email: email,
		Password: string(hash),
		AvatarUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ4YreOWfDX3kK-QLAbAL4ufCPc84ol2MA8Xg&s",
	}
}