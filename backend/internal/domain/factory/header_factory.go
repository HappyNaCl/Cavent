package factory

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/google/uuid"
)

type HeaderFactory struct {}

func NewHeaderFactory() *HeaderFactory {
	return &HeaderFactory{}
}

func (f *HeaderFactory) CreateTransactionHeader(userId string, eventId uuid.UUID) *model.TransactionHeader {
	return &model.TransactionHeader{
		Id:      uuid.New(),
		UserId:  userId,
		EventId: eventId,
	}
}