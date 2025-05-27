package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"gorm.io/gorm"
)

type AnalyticsRepository struct {
	db *gorm.DB
}

// CreateEventView implements repo.AnalyticsRepository.
func (a *AnalyticsRepository) CreateEventView(view *model.EventView) error {
	if err := a.db.Create(view).Error; err != nil {
		return err
	}
	return nil
}

// GetEventViews implements repo.AnalyticsRepository.
func (a *AnalyticsRepository) GetEventViews(eventID string) (int, error) {
	panic("unimplemented")
}

func NewAnalyticsRepository(db *gorm.DB) repo.AnalyticsRepository {
	return &AnalyticsRepository{
		db: db,
	}
}
