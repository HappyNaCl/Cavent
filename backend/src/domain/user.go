package domain

import "gorm.io/gorm"

// Make a struct User that  represents the user model
type User struct {
	gorm.Model
	Provider string `json:"provider" gorm:"index"`
	ProviderID string `json:"provider_id" gorm:"uniqueIndex"`
	Email string `json:"email"`
	Name string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}