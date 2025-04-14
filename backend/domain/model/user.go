package model

type User struct {
	Id string `json:"id" gorm:"primaryKey"`
	Provider string `json:"provider"`
	Email string `json:"email" gorm:"unique; not null"`
	Name string `json:"name" gorm:"not null"`
	Password string `json:"-" gorm:"not null"`
	AvatarUrl string `json:"avatarUrl gorm:not null"`
	FirstTimeLogin bool `json:"firstTimeLogin" gorm:"default:true"`
	Description string `json:"description" gorm:"size:500"`
	CreatedAt int64 `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt int64 `json:"updatedAt" gorm:"autoUpdateTime"`
	Tags []Tag `json:"tags" gorm:"many2many:user_interests;joinForeignKey:UserId;joinReferences:TagId"`
}