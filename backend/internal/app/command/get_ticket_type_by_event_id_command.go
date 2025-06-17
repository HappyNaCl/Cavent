package command

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type GetTicketTypeByEventIdCommand struct {
	EventId string `json:"eventId"`
}

type GetTicketTypeByEventIdCommandResult struct {
	TicketTypes []*common.TicketTypeResult `json:"ticketTypes"`
}