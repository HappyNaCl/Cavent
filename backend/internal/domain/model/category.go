package model

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
    Id              uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
    CategoryTypeId  uuid.UUID `json:"categoryTypeId" gorm:"type:uuid;not null"`
    CategoryType    CategoryType
    Name            string    `json:"name" gorm:"unique;not null"`
	CreatedAt       time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updatedAt" gorm:"autoUpdateTime"`

    Users      []User     `gorm:"many2many:user_interests"`
    Events     []Event   `gorm:"many2many:event_categories;"`
}