package persistence

import (
	"github.com/HappyNaCl/Cavent/src/domain"
	"github.com/HappyNaCl/Cavent/src/domain/factory"
	"github.com/HappyNaCl/Cavent/src/domain/repository"
	"github.com/markbates/goth"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct{
	Conn *gorm.DB
}

func UserRepository(conn *gorm.DB) repository.UserRepository{
	return &UserRepositoryImpl{Conn: conn}
}

func (repo *UserRepositoryImpl) FindByProviderID(providerID string) (*domain.User, error){
	var user domain.User
	err := repo.Conn.Preload("User").Where("provider_id = ?", providerID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) FindByEmail(email string) (*domain.User, error){
	var user domain.User
	err := repo.Conn.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) RegisterUser(user *domain.User) (*domain.User, error){
	err := repo.Conn.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepositoryImpl) RegisterOrLoginOauthUser(gothUser goth.User, provider string) (*domain.User, error){
	var user domain.User

	result := repo.Conn.Where("provider_id = ? AND provider = ?", gothUser.UserID, provider).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			user := factory.UserFactory().GetOAuthUser(provider, gothUser.UserID, gothUser.Name, gothUser.Email, "", gothUser.AvatarURL)
			repo.Conn.Create(&user)
		} else {
			return nil, result.Error
		}
	}

	return &user, nil
}