package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CampusGormRepo struct {
	db *gorm.DB
}

// GetCampusById implements repo.CampusRepository.
func (c *CampusGormRepo) GetCampusById(campusId uuid.UUID) (*model.Campus, error) {
	var campus model.Campus

	err := c.db.First(&campus, "id = ?", campusId).Error
	if err != nil {
		return nil, err
	}

	return &campus, nil
}

func (c *CampusGormRepo) FindCampusByInviteCode(inviteCode string) (*model.Campus, error) {
	var campus model.Campus
	err := c.db.Where("invite_code = ?", inviteCode).First(&campus).Error
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
		db: db,
	}
}
