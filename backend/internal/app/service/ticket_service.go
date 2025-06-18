package service

import (
	"context"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb"
	"github.com/google/uuid"
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

func (s *TicketService) GetUserTickets(ctx context.Context, com *command.GetUserTicketsCommand) (*command.GetUserTicketsCommandResponse, error) {
	tickets, err := s.ticketRepository.GetUserTickets(com.UserId)
	if err != nil {
		return nil, err
	}

	eventMap := make(map[uuid.UUID]*common.UserTicketResult)

	for _, ticket := range tickets {
		eventID := ticket.TicketType.Event.Id

		if existing, found := eventMap[eventID]; found {
			existing.Tickets = append(existing.Tickets, common.TicketResult{
				Id:   ticket.Id,
				Name: ticket.TicketType.Name,
			})
		} else {
			eventMap[eventID] = &common.UserTicketResult{
				EventId:    eventID,
				EventTitle: ticket.TicketType.Event.Title,
				StartTime:  ticket.TicketType.Event.StartTime.Format("2006-01-02 15:04:05"),
				EndTime:    ticket.TicketType.Event.EndTime.Format("2006-01-02 15:04:05"),
				Tickets: []common.TicketResult{
					{
						Id:   ticket.Id,
						Name: ticket.TicketType.Name,
					},
				},
			}
		}
	}

	var userTickets []*common.UserTicketResult
	for _, userTicket := range eventMap {
		userTickets = append(userTickets, userTicket)
	}

	return &command.GetUserTicketsCommandResponse{
		Tickets: userTickets,
	}, nil
}