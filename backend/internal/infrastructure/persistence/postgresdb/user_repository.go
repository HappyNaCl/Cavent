package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/errors"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserGormRepo struct {
	db *gorm.DB
}

// GetUserProfile implements repo.UserRepository.
func (u *UserGormRepo) GetUserProfile(userId string) (*model.User, error) {
	var user model.User
	err := u.db.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUserProfile implements repo.UserRepository.
func (u *UserGormRepo) UpdateUserProfile(user *model.User) (*model.User, error) {
	var oldUser model.User
	err := u.db.Where("id = ?", user.Id).First(&oldUser).Error
	if err != nil {
		return nil, err
	}

	oldUser.Name = user.Name
	oldUser.AvatarUrl = user.AvatarUrl
	oldUser.Description = user.Description
	oldUser.PhoneNumber = user.PhoneNumber
	oldUser.Address = user.Address

	if err := u.db.Save(&oldUser).Error; err != nil {
		return nil, err
	}
	return &oldUser, nil
}

// UpdateUserCampus implements repo.UserRepository.
func (u *UserGormRepo) UpdateUserCampus(userId string, campusId uuid.UUID) (*model.User, error) {
	var user model.User
	err := u.db.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}

	if user.CampusId != nil && *user.CampusId == campusId {
		return &user, nil
	}

	user.CampusId = &campusId
	if err := u.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByEmail implements repo.UserRepository.
func (u *UserGormRepo) GetBriefUser(userId string) (*model.User, error) {
	var user model.User
	err := u.db.Model(&model.User{}).Select("id, campus_id, provider, name, avatar_url, email, first_time_login, role").Where("id = ?", userId).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserGormRepo) GetUserInterests(userId string) ([]*model.Category, error) {
	var user model.User
	err := u.db.Where("id = ?", userId).Preload("Interests").First(&user).Error
	if err != nil {
		return nil, err
	}

	interests := make([]*model.Category, 0)
	for _, interest := range user.Interests {
		interests = append(interests, &interest)
	}

	return interests, nil
}

func (u *UserGormRepo) UpdateUserInterests(userId string, interestIds []string) error {
	var user model.User
	err := u.db.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return err
	}
	var interests []model.Category
	if err := u.db.Where("id IN ?", interestIds).Find(&interests).Error; err != nil {
		return err
	}

	if err := u.db.Model(&user).Association("Interests").Replace(interests); err != nil {
		return err
	}

	user.FirstTimeLogin = false
	if err := u.db.Model(&user).Update("first_time_login", false).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserGormRepo) GetCampusId(userId string) (*uuid.UUID, error) {
	var user model.User
	err := u.db.Model(&model.User{}).Select("campus_id").Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}

	if user.CampusId == nil {
		return nil, errors.ErrUserNotInCampus
	}

	return user.CampusId, nil
}

func NewUserGormRepo(db *gorm.DB) repo.UserRepository {
	return &UserGormRepo{
		db: db,
	}
}
