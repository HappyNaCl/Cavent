package factory

import (
	"github.com/HappyNaCl/Cavent/src/config"
	"github.com/HappyNaCl/Cavent/src/domain"
	"github.com/google/uuid"
)

type UserFactoryInterface interface {
	GetOAuthUser(provider string, providerId string, name string, email string, password string, avatarUrl string) *domain.User
	GetUser(name string, email string, password string) *domain.User
}

type UserFactoryImpl struct {}

func (u *UserFactoryImpl) GetOAuthUser(provider string, providerId string, name string, email string, password string, avatarUrl string) *domain.User {
	hash, err := config.HashPassword(password)

	if err != nil {
		return nil
	}

	return &domain.User{
		Provider: provider,
		ProviderID: providerId,
		Name: name,
		Email: email,
		Password: hash,
		AvatarUrl: avatarUrl,
	}
}

func (u *UserFactoryImpl) GetUser(name string, email string, password string) *domain.User {
	hash, err := config.HashPassword(password)

	if err != nil {
		return nil
	}

	return &domain.User{
		Provider: "credential",
		ProviderID: uuid.NewString(),
		Name: name,
		Email: email,
		Password: hash,
		AvatarUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ4YreOWfDX3kK-QLAbAL4ufCPc84ol2MA8Xg&s",
	}
}

func UserFactory() UserFactoryInterface {
	return &UserFactoryImpl{}
}
