package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"gorm.io/gorm"
)

type CampusGormRepo struct {
	db    *gorm.DB
}

func (c *CampusGormRepo) FindCampusByInviteCode(inviteCode string) (*model.Campus, error) {
	var campus model.Campus

	err := c.db.Find(&campus).Where("inviteCode = ?", inviteCode).Error
	if err != nil {
		return nil, err
	}

	return &campus, nil
}

func (c *CampusGormRepo) GetAllCampus() ([]*model.Campus, error) {
	var campuses []*model.Campus
	
	err := c.db.Find(&campuses).Error
	if err != nil {
		return nil, err
	}

	return campuses, nil
}

func NewCampusGormRepo(db *gorm.DB) repo.CampusRepository {
	return &CampusGormRepo{
		db:    db,
	}
}
