package persistence

import (
	"strings"

	"github.com/HappyNaCl/Cavent/backend/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/domain/model"
	"github.com/HappyNaCl/Cavent/backend/domain/repository"
	"github.com/markbates/goth"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct{
	Conn *gorm.DB
}

func UserRepository(conn *gorm.DB) repository.UserRepository{
	return &UserRepositoryImpl{Conn: conn}
}

func (repo *UserRepositoryImpl) FindByProviderID(providerID string) (*model.User, error){
	var user model.User
	err := repo.Conn.Preload("User").Where("id = ?", providerID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) FindByEmail(email string) (*model.User, error){
	var user model.User
	err := repo.Conn.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) RegisterUser(user *model.User) (*model.User, error){
	err := repo.Conn.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepositoryImpl) RegisterOrLoginOauthUser(gothUser goth.User, provider string) (*model.User, error){
	var user model.User

	result := repo.Conn.Where("provider_id = ? AND provider = ?", gothUser.UserID, provider).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			user := factory.UserFactory().GetOAuthUser(provider, gothUser.UserID, strings.Split(gothUser.Email, "@")[0], gothUser.Email, "", gothUser.AvatarURL)
			repo.Conn.Create(&user)
		} else {
			return nil, result.Error
		}
	}

	return &user, nil
}

func (repo *UserRepositoryImpl) UpdateInterest(userId string, preferences []string) error {
	var user model.User
	if err := repo.Conn.Where("id = ?", userId).First(&user).Error; err != nil {
		return err
	}

	var tags []model.Tag
	if err := repo.Conn.Where("id IN ?", preferences).Find(&tags).Error; err != nil {
		return err
	}

	if err := repo.Conn.Model(&user).Association("Tags").Replace(&tags); err != nil {
		return err
	}
	
	user.FirstTimeLogin = false

	if err := repo.Conn.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (repo *UserRepositoryImpl) GetUserTag(userId string) ([]model.Tag, error) {
	var user model.User
	if err := repo.Conn.Preload("Tags").Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return user.Tags, nil
}