package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserGormRepo struct {
	db    *gorm.DB
}

// FindByEmail implements repo.UserRepository.
func (u *UserGormRepo) GetBriefUser(userId uuid.UUID) (*model.User, error) {
	var user model.User
	err := u.db.Model(&model.User{}).Select("id, email, first_time_login, role").Where("id = ?", userId).First(&user).Error; 
	
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserGormRepo(db *gorm.DB) repo.UserRepository {
	return &UserGormRepo{
		db:    db,
	}
}
