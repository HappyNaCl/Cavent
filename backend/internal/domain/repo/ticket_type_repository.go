package repo

import "github.com/HappyNaCl/Cavent/backend/internal/domain/model"

type TicketTypeRepository interface {
	GetTicketTypeByEventId(eventId string) ([]*model.TicketType, error)
}