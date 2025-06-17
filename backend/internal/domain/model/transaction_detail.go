package model

import (
	"time"

	"github.com/google/uuid"
)

type TransactionDetail struct {
	Id      	  uuid.UUID  `gorm:"type:uuid;primaryKey"`
	TicketTypeId  uuid.UUID  `gorm:"not null"`
	Quantity      int        `gorm:"not null"`
	CreatedAt     time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
}