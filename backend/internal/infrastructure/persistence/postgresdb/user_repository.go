package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"gorm.io/gorm"
)

type UserGormRepo struct {
	db    *gorm.DB
}

// FindByEmail implements repo.UserRepository.
func (u *UserGormRepo) GetBriefUser(userId string) (*model.User, error) {
	var user model.User
	err := u.db.Model(&model.User{}).Select("id, campus_id, provider, name, avatar_url, email, first_time_login, role").Where("id = ?", userId).First(&user).Error; 
	
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
