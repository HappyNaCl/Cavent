package service

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb"
	"gorm.io/gorm"
)

type AnalyticsService struct {
	repository repo.AnalyticsRepository
}

func NewAnalyticsService(db *gorm.DB) *AnalyticsService {
	return &AnalyticsService{
		repository: postgresdb.NewAnalyticsGormRepo(db),
	}
}

func (s *AnalyticsService) CreateEventView(com *command.CreateEventViewCommand) (*command.CreateEventViewCommandResult, error) {
	factory := factory.NewEventViewFactory()
	eventView := factory.CreateEventViewCommand(com.EventID, com.UserID)

	if err := s.repository.CreateEventView(eventView); err != nil {
		return nil, err
	}

	return &command.CreateEventViewCommandResult{}, nil
}