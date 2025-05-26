package repo

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
)

type UserRepository interface {
	GetBriefUser(userId string) (*model.User, error)
	GetUserInterests(userId string) ([]*model.Category, error)
	UpdateUserInterests(userId string, interestIds []string) error
}