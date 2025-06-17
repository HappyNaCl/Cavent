package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"gorm.io/gorm"
)

type TicketTypeGormRepo struct {
	db *gorm.DB
}

func (t *TicketTypeGormRepo) GetTicketTypeByEventId(eventId string) ([]*model.TicketType, error) {
	var ticketTypes []*model.TicketType
	if err := t.db.Where("event_id = ?", eventId).Find(&ticketTypes).Error; err != nil {
		return nil, err
	}
	return ticketTypes, nil
}

func NewTicketTypeGormRepo(db *gorm.DB) repo.TicketTypeRepository {
	return &TicketTypeGormRepo{
		db: db,
	}
}
