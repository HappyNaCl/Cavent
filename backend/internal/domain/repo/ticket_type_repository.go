package repo

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/google/uuid"
)

type TicketTypeRepository interface {
	GetTicketTypeByEventId(eventId string) ([]*model.TicketType, error)
	ReduceTicketTypeQuantity(ticketTypeId uuid.UUID) error
}