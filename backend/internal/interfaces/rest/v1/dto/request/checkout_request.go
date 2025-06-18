package request

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/google/uuid"
)

type CheckoutTicketDTO struct {
	Id 		 string  `json:"ticketId" form:"ticketId" binding:"required"`
	Quantity int     `json:"quantity" form:"quantity" binding:"required"`
}

type CheckoutRequest struct {
	EventId string `json:"eventId" form:"eventId" binding:"required"`
	Tickets  []CheckoutTicketDTO `json:"tickets" form:"tickets" binding:"required"`   
}

func (r *CheckoutRequest) ToCommand() (*command.CheckoutCommand , error){
	ticketResult := make([]*common.CheckoutTicketResult, len(r.Tickets))
	for i, ticket := range r.Tickets {

		ticketUuid, err := uuid.Parse(ticket.Id)
		if err != nil {

			return nil, err
		} 

		ticketResult[i] = &common.CheckoutTicketResult{
			Id:       ticketUuid,
			Quantity: ticket.Quantity,
		}
	}

	return &command.CheckoutCommand{
		Ticket: ticketResult,
	}, nil
}