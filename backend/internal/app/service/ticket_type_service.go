package service

import (
	"context"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/mapper"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb"
	"gorm.io/gorm"
)

type TicketTypeService struct {
	ticketTypeRepo repo.TicketTypeRepository
}

func NewTicketTypeService(db *gorm.DB) *TicketTypeService {
	return &TicketTypeService{
		ticketTypeRepo: postgresdb.NewTicketTypeGormRepo(db),
	}
}

func (tts *TicketTypeService) GetTicketTypeByEventId(ctx context.Context, com *command.GetTicketTypeByEventIdCommand) (*command.GetTicketTypeByEventIdCommandResult, error){
	ticketTypes, err := tts.ticketTypeRepo.GetTicketTypeByEventId(com.EventId)
	if err != nil {
		return nil, err
	}

	return &command.GetTicketTypeByEventIdCommandResult{
		TicketTypes: mapper.NewTicketTypeResultFromTicketType(ticketTypes),
	}, nil
}