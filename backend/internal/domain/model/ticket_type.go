package model

import (
	"time"

	"github.com/google/uuid"
)

type TicketType struct {
    Id        uuid.UUID `gorm:"type:uuid;primaryKey"`
    EventId   uuid.UUID
    Event     Event
    Name      string    `json:"name" gorm:"unique;not null"`
    Price     float64   `gorm:"type:decimal(10,2)"`
    Quantity  int       `json:"quantity"`
    Sold      int       `json:"sold" gorm:"default:0"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
    Tickets   []Ticket  `gorm:"foreignKey:TicketTypeId"`
}