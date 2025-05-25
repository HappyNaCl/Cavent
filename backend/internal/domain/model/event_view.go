package model

import (
	"time"

	"github.com/google/uuid"
)

type EventView struct {
    Id        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
    EventId   uuid.UUID `json:"eventId" gorm:"type:uuid;index"`
    Event     Event
    UserId    uuid.UUID `json:"userId" gorm:"type:uuid"`
    User      User
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}