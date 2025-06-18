package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/google/uuid"
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

func (t *TicketTypeGormRepo) ReduceTicketTypeQuantity(ticketTypeId uuid.UUID) error {
	if err := t.db.Model(&model.TicketType{}).Where("id = ?", ticketTypeId).UpdateColumn("quantity", gorm.Expr("quantity - 1")).Error; err != nil {
		return err
	}
	return nil
}

func (t *TicketTypeGormRepo) IsTicketAvailable(ticketTypeId uuid.UUID) (bool, error) {
	var count int64
	if err := t.db.Model(&model.TicketType{}).Where("id = ? AND quantity > 0 AND sold < quantity", ticketTypeId).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func NewTicketTypeGormRepo(db *gorm.DB) repo.TicketTypeRepository {
	return &TicketTypeGormRepo{
		db: db,
	}
}
