package model

import (
	"time"

	"github.com/google/uuid"
)

type TransactionDetail struct {
	Id      	  uuid.UUID  `gorm:"type:uuid;primaryKey"`
	TransactionId uuid.UUID  `gorm:"not null"`
	TicketTypeId  uuid.UUID  `gorm:"not null"`
	Quantity      int        `gorm:"not null"`

	TicketType        TicketType        `gorm:"foreignKey:TicketTypeId;references:Id"`
	TransactionHeader TransactionHeader `gorm:"foreignKey:TransactionId;references:Id"`

	CreatedAt     time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
}