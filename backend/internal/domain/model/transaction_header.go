package model

import (
	"time"

	"github.com/google/uuid"
)

type TransactionHeader struct {
	Id      uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserId  string    `gorm:"not null"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}