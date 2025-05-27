package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"gorm.io/gorm"
)

type EventGormRepo struct {
	db *gorm.DB
}

// CreateEvent implements repo.EventRepository.
func (e *EventGormRepo) CreateEvent(event *model.Event) (*model.Event, error) {
	panic("unimplemented")
}

func NewEventGormRepo(db *gorm.DB) repo.EventRepository {
	return &EventGormRepo{
		db: db,
	}
}
