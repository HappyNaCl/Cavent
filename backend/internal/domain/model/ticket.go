package model

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
    Id            uuid.UUID `gorm:"type:uuid;primaryKey"`
    UserId        uuid.UUID
    User          User
    TicketTypeId  uuid.UUID
    TicketType    TicketType
    CreatedAt     time.Time
    UpdatedAt     time.Time
}