package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"gorm.io/gorm"
)

type CampusRepository struct {
	db    *gorm.DB
}

func (c *CampusRepository) FindCampusByInviteCode(inviteCode string) (*model.Campus, error) {
	var campus model.Campus

	err := c.db.Find(&campus).Where("inviteCode = ?", inviteCode).Error
	if err != nil {
		return nil, err
	}

	return &campus, nil
}

func NewCampusGormRepo(db *gorm.DB) repo.CampusRepository {
	return &CampusRepository{
		db:    db,
	}
}
