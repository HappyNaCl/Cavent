package repo

import "github.com/HappyNaCl/Cavent/backend/internal/domain/model"

type AnalyticsRepository interface {
	CreateEventView(view *model.EventView) error
	GetEventViews(eventID string) (int, error)
}