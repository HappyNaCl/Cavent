package model

import (
	"time"

	"github.com/google/uuid"
)

type CategoryType struct {
    Id        uuid.UUID `gorm:"type:uuid;primaryKey"`
    Name      string    `json:"name" gorm:"unique;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
    Categories []Category `gorm:"foreignKey:CategoryTypeId"`
}