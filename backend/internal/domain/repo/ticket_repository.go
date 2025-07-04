package repo

import "github.com/HappyNaCl/Cavent/backend/internal/domain/model"

type TicketRepository interface {
	GetUserTickets(userId string) ([]*model.Ticket, error)
	CreateUserTicket(ticket *model.Ticket) (*model.Ticket, error)
}