package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"gorm.io/gorm"
)

type TicketGormRepo struct {
	db *gorm.DB
}

func NewTicketGormRepo(db *gorm.DB) repo.TicketRepository {
	return &TicketGormRepo{
		db: db,
	}
}