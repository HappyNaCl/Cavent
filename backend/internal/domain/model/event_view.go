package model

import (
	"time"

	"github.com/google/uuid"
)

type EventView struct {
    Id        uuid.UUID `gorm:"type:uuid;primaryKey"`
    EventId   uuid.UUID `gorm:"type:uuid;index"`
    Event     Event
    UserId    uuid.UUID `gorm:"type:uuid"`
    User      User
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}