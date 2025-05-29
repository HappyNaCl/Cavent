package repo

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/google/uuid"
)

type EventRepository interface {
	CreateEvent(event *model.Event) (*model.Event, error)
	GetEvents(limit int) ([]*model.Event, error)
	GetEventsByCategories(categories []uuid.UUID, limit int) ([]*model.Event, error) 
}