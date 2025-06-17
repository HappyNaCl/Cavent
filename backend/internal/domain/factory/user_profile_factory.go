package factory

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
)

type UserProfileFactory struct{}

func NewUserProfileFactory() *UserProfileFactory {
	return &UserProfileFactory{}
}

func (f *UserProfileFactory) CreateUserProfileResult(id, name string, description, phoneNumber, address *string) *model.User {
	result :=  &model.User{
		Id:          id,
		Name:        name,
	}

	if description != nil {
		result.Description = description
	}
	if phoneNumber != nil {
		result.PhoneNumber = phoneNumber
	}
	if address != nil {
		result.Address = address
	}

	return result
}