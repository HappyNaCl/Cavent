package postgresdb

import (
	"time"

	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb/paginate"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type EventGormRepo struct {
	db *gorm.DB
}

// SearchEvents implements repo.EventRepository.
func (e *EventGormRepo) SearchEvents(query string) ([]*model.Event, error) {
	var events []*model.Event
	err := e.db.Raw(`
		SELECT id, title, location, start_time, similarity(lower(title), ?) AS score
		FROM events
		ORDER BY score DESC
		LIMIT 5
	`, query).Scan(&events).Error
	if err != nil {
		return nil, err
	}

	zap.L().Sugar().Infof("Search query: %s, found %d events", query, len(events))
	return events, nil
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

func (e *EventGormRepo) GetAllEvents(pagination paginate.Pagination) (*paginate.Pagination, error){
	var events []*model.Event

	query := e.db.Preload("TicketTypes").Preload("Campus").Preload("Category")

	for i, filter := range pagination.Filter{
		zap.L().Sugar().Infof("Filter %d: %s with args %v", i, filter, pagination.FilterArgs[i])
		args := pagination.FilterArgs[i]
		query = query.Where(filter, args...)
	}

	query = query.Scopes(paginate.Paginate(events, &pagination, e.db)).
			Where("start_time > ?", time.Now())


	err := query.Find(&events).Error

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

func (e *EventGormRepo) GetUserFavoritedEvent(userId string) ([]*model.Event, error) {
    var events []*model.Event
    err := e.db.Model(&model.Event{}).
        Joins("JOIN user_favorites ON user_favorites.event_id = events.id").
        Where("user_favorites.user_id = ?", userId).
        Preload("TicketTypes").
        Preload("Campus").
        Preload("Category").
        Order("events.start_time ASC").
        Find(&events).Error

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
