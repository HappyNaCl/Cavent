package factory

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/google/uuid"
)

type TicketFactory struct {

}

func NewTicketFactory() *TicketFactory {
	return &TicketFactory{}
}

func (f *TicketFactory) CreateTicket(ticketTypeId uuid.UUID, userId string) *model.Ticket {
	return &model.Ticket{
		Id: 		  uuid.New(),
		TicketTypeId: ticketTypeId,
		UserId: 	  userId,
	}
}