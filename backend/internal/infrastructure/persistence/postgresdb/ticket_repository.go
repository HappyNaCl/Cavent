package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"go.uber.org/zap"
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
	var tickets []*model.Ticket
	zap.L().Sugar().Debugf("Fetching tickets for user: %s", userId)
	if err := t.db.Where("user_id = ?", userId).Preload("TicketType").Preload("TicketType.Event").Find(&tickets).Error; err != nil {
		return nil, err
	}

	if len(tickets) == 0 {
		zap.L().Sugar().Debugf("No tickets found for user: %s", userId)
	}
	return tickets, nil
}

func NewTicketGormRepo(db *gorm.DB) repo.TicketRepository {
	return &TicketGormRepo{
		db: db,
	}
}
