package model

import (
	"time"

	"github.com/google/uuid"
)

type EventView struct {
    Id        uuid.UUID `gorm:"type:uuid;primaryKey"`
    EventId   uuid.UUID
    Event     Event
    UserId    uuid.UUID
    User      User
    ViewedAt  time.Time
}