package repo

import "github.com/HappyNaCl/Cavent/backend/internal/domain/model"

type EventRepository interface {
	CreateEvent(event *model.Event) (*model.Event, error)
	GetEvents(limit int) ([]*model.Event, error)
}