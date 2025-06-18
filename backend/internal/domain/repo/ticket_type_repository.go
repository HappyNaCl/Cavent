package repo

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/google/uuid"
)

type TicketTypeRepository interface {
	IsTicketAvailable(ticketTypeId uuid.UUID) (bool, error)
	GetTicketTypeByEventId(eventId string) ([]*model.TicketType, error)
	ReduceTicketTypeQuantity(ticketTypeId uuid.UUID) error
}