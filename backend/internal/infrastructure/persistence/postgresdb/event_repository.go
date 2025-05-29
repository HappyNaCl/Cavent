package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"gorm.io/gorm"
)

type EventGormRepo struct {
	db *gorm.DB
}

// GetEvents implements repo.EventRepository.
func (e *EventGormRepo) GetEvents(limit int) ([]*model.Event, error) {
	var events []*model.Event
	err := e.db.Preload("TicketTypes").Preload("Campus").Preload("Category").Order("start_time ASC").Limit(limit).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

// CreateEvent implements repo.EventRepository.
func (e *EventGormRepo) CreateEvent(event *model.Event) (*model.Event, error) {
	tx := e.db.Begin()

	if err := tx.Create(&event).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for i := range event.TicketTypes {
		event.TicketTypes[i].EventId = event.Id
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return event, nil
}

func NewEventGormRepo(db *gorm.DB) repo.EventRepository {
	return &EventGormRepo{
		db: db,
	}
}
