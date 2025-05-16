package model

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
    Id         uuid.UUID `gorm:"type:uuid;primaryKey"`
    CategoryTypeId  uuid.UUID
    CategoryType    CategoryType
    Name       string
    CreatedAt  time.Time
    UpdatedAt  time.Time

    Users      []User     `gorm:"many2many:user_interests"`
    Events     []Event    `gorm:"foreignKey:CategoryId"`
}