package repo

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb/paginate"
	"github.com/google/uuid"
)

type EventRepository interface {
	CreateEvent(event *model.Event) (*model.Event, error)
	GetEvents(limit int) ([]*model.Event, error)
	GetEventsByCategories(categories []uuid.UUID, limit int) ([]*model.Event, error) 
	GetEventByID(eventID uuid.UUID) (*model.Event, error)
	GetCampusEvents(campusID uuid.UUID, pagination paginate.Pagination) (*paginate.Pagination, error)
}