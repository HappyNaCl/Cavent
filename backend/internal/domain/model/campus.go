package model

import (
	"time"

	"github.com/google/uuid"
)

type Campus struct {
    Id          uuid.UUID `gorm:"type:uuid;primaryKey"`
    Name        string
    LogoUrl     string
    Description string
    CreatedAt   time.Time
    UpdatedAt   time.Time
    Users       []User    `gorm:"foreignKey:CampusId"`
    Events      []Event   `gorm:"foreignKey:CampusId"`
}
