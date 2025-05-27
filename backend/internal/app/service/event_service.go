package service

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb"
	"gorm.io/gorm"
)

type EventService struct {
	eventRepo repo.EventRepository
}

func NewEventService(db *gorm.DB) *EventService {
	return &EventService{
		eventRepo: postgresdb.NewEventGormRepo(db),
	}
}
