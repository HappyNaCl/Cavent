package command

import "github.com/google/uuid"

type CreateTicketCommand struct {
	UserId string
	TicketId uuid.UUID
}

type CreateTicketCommandResult struct {}