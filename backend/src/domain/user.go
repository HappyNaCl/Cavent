package domain

import "gorm.io/gorm"

// Make a struct User that  represents the user model
type User struct {
	gorm.Model
	Provider string `json:"provider" gorm:"index"`
	ProviderID string `json:"provider_id" gorm:"uniqueIndex"`
	Email string `json:"email" gorm:"unique"`
	Name string `json:"name"`
	Password string `json:"-"`
	AvatarUrl string `json:"avatar_url"`
}