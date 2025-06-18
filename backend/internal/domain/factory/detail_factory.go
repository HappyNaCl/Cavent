package factory

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/google/uuid"
)


type DetailFactory struct{}

func NewDetailFactory() *DetailFactory {
	return &DetailFactory{}
}

func (f *DetailFactory) CreateTransactionDetail(transactionId uuid.UUID, ticketId uuid.UUID, quantity int) model.TransactionDetail {
	return model.TransactionDetail{
		Id:             uuid.New(),
		TransactionId:  transactionId,
		TicketTypeId:   ticketId,
		Quantity:       quantity,
	}
}