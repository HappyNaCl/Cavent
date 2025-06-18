package command

import "github.com/google/uuid"

type ReduceTicketTypeQuantityCommand struct {
	TicketTypeId  uuid.UUID `json:"ticketTypeId"`
}

type ReduceTicketTypeQuantityCommandResult struct {}