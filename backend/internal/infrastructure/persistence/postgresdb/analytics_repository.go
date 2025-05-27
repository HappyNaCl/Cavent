package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"gorm.io/gorm"
)

type AnalyticsGormRepo struct {
	db *gorm.DB
}

// CreateEventView implements repo.AnalyticsGormRepo.
func (a *AnalyticsGormRepo) CreateEventView(view *model.EventView) error {
	if err := a.db.Create(view).Error; err != nil {
		return err
	}
	return nil
}

// GetEventViews implements repo.AnalyticsGormRepo.
func (a *AnalyticsGormRepo) GetEventViews(eventID string) (int, error) {
	panic("unimplemented")
}

func NewAnalyticsGormRepo(db *gorm.DB) repo.AnalyticsRepository {
	return &AnalyticsGormRepo{
		db: db,
	}
}
