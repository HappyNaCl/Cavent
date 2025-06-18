package service

import (
	"context"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb"
	"gorm.io/gorm"
)

type TicketService struct {
	ticketRepository     repo.TicketRepository
	ticketTypeRepository repo.TicketTypeRepository
}

func NewTicketService(db *gorm.DB) *TicketService {
	return &TicketService{
		ticketRepository: postgresdb.NewTicketGormRepo(db),
		ticketTypeRepository: postgresdb.NewTicketTypeGormRepo(db),
	}
}

func (s *TicketService) CreateTicket(ctx context.Context, com *command.CreateTicketCommand) (*command.CreateTicketCommandResult, error) {
	factory := factory.NewTicketFactory()
	ticket := factory.CreateTicket(com.TicketId, com.UserId)

	if _, err := s.ticketRepository.CreateUserTicket(ticket); err != nil {
		return nil, err
	}

	if err := s.ticketTypeRepository.ReduceTicketTypeQuantity(com.TicketId); err != nil {
		return nil, err
	}

	return &command.CreateTicketCommandResult{
	}, nil
}