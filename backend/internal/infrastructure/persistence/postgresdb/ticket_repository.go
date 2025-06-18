package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"gorm.io/gorm"
)

type TicketGormRepo struct {
	db *gorm.DB
}

// CreateUserTickets implements repo.TicketRepository.
func (t *TicketGormRepo) CreateUserTicket(ticket *model.Ticket) (*model.Ticket, error) {
	if err := t.db.Create(ticket).Error; err != nil {
		return nil, err
	}
	return ticket, nil
}

// GetUserTickets implements repo.TicketRepository.
func (t *TicketGormRepo) GetUserTickets(userId string) ([]*model.Ticket, error) {
	panic("unimplemented")
}

func NewTicketGormRepo(db *gorm.DB) repo.TicketRepository {
	return &TicketGormRepo{
		db: db,
	}
}
