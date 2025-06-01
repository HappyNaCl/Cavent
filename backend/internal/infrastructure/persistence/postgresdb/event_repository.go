package postgresdb

import (
	"time"

	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb/paginate"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventGormRepo struct {
	db *gorm.DB
}

// GetCampusEvents implements repo.EventRepository.
func (e *EventGormRepo) GetCampusEvents(campusID uuid.UUID, pagination paginate.Pagination) (*paginate.Pagination, error) {
	var events []*model.Event
	err := e.db.Preload("TicketTypes").Preload("Campus").Preload("Category").
		Scopes(paginate.Paginate(events, &pagination, e.db)).
		Where("campus_id = ?", campusID).
		Where("start_time > ?", time.Now()).
		Order("start_time ASC").Find(&events).Error
	
	if err != nil {
		return nil, err
	}

	pagination.Rows = events
	return &pagination, nil
}

// GetEventByID implements repo.EventRepository.
func (e *EventGormRepo) GetEventByID(eventID uuid.UUID) (*model.Event, error) {
	var event *model.Event
	err := e.db.Preload("TicketTypes").Preload("Campus").Preload("Category").
		Where("id = ?", eventID).First(&event).Error
	if err != nil {
		return nil, err
	}
	return event, nil
}

// GetEvents implements repo.EventRepository.
func (e *EventGormRepo) GetEvents(limit int) ([]*model.Event, error) {
	var events []*model.Event
	err := e.db.Preload("TicketTypes").Preload("Campus").Preload("Category").
		Where("start_time > ?", time.Now()).Order("start_time ASC").Limit(limit).Find(&events).Error
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

func (e *EventGormRepo) GetEventsByCategories(categories []uuid.UUID, limit int) ([]*model.Event, error) {
	var events []*model.Event
	err := e.db.Preload("TicketTypes").Preload("Campus").Preload("Category").
		Where("category_id IN ?", categories).
		Where("start_time > ?", time.Now()).
		Order("start_time ASC").Limit(limit).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func NewEventGormRepo(db *gorm.DB) repo.EventRepository {
	return &EventGormRepo{
		db: db,
	}
}
