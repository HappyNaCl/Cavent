package repo

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
)

type TransactionRepository interface {
	CreateTransaction(transaction *model.TransactionHeader) error
}