package model

import (
	"time"

	"github.com/google/uuid"
)

type Campus struct {
    Id          uuid.UUID `gorm:"type:uuid;primaryKey"`
    Name        string  `json:"name"`
    LogoUrl     string  `json:"logoUrl"`
    Description string  `json:"description"`
    InviteCode  string  `json:"inviteCode" gorm:"unique;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
    Users       []User    `gorm:"foreignKey:CampusId"`
    Events      []Event   `gorm:"foreignKey:CampusId"`
}
