package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"gorm.io/gorm"
)

type TransactionGormRepo struct {
	db *gorm.DB
}

// CreateTransaction implements repo.TransactionRepository.
func (t *TransactionGormRepo) CreateTransaction(userId string, eventId string, ticketTypes []*model.TicketType) error {
	panic("unimplemented")
}

func NewTransactionGormRepo(db *gorm.DB) repo.TransactionRepository {
	return &TransactionGormRepo{
		db: db,
	}
}
