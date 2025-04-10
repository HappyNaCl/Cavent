package factory

import (
	"github.com/HappyNaCl/Cavent/src/domain"
	"github.com/google/uuid"
)

type UserFactoryInterface interface {
	GetOAuthUser(provider string, providerId string, name string, email string, password string, avatarUrl string) *domain.User
	GetUser(name string, email string, password string) *domain.User
}

type UserFactoryImpl struct {}

func (u *UserFactoryImpl) GetOAuthUser(provider string, providerId string, name string, email string, password string, avatarUrl string) *domain.User {
	return &domain.User{
		Provider: provider,
		ProviderID: providerId,
		Name: name,
		Email: email,
		Password: password,
		AvatarUrl: avatarUrl,
	}
}

func (u *UserFactoryImpl) GetUser(name string, email string, password string) *domain.User {
	return &domain.User{
		Provider: "credential",
		ProviderID: uuid.NewString(),
		Name: name,
		Email: email,
		Password: password,
		
	}
}

func UserFactory() UserFactoryInterface {
	return &UserFactoryImpl{}
}
