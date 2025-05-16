package model

import (
	"time"

	"github.com/google/uuid"
)

type TicketType struct {
    Id        uuid.UUID `gorm:"type:uuid;primaryKey"`
    EventId   uuid.UUID
    Event     Event
    Name      string
    Price     float64   `gorm:"type:decimal(10,2)"`
    Quantity  int
    CreatedAt time.Time
    UpdatedAt time.Time
    Tickets   []Ticket  `gorm:"foreignKey:TicketTypeId"`
}