package repo

import "github.com/HappyNaCl/Cavent/backend/internal/domain/model"

type TransactionRepository interface {
	CreateTransaction(userId string, eventId string, ticketTypes []*model.TicketType) error
}