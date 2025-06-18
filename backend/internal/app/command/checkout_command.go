package command

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/google/uuid"
)

type CheckoutCommand struct {
	UserId string
	EventId uuid.UUID
	Ticket []*common.CheckoutTicketResult
}

type CheckoutCommandResult struct {}