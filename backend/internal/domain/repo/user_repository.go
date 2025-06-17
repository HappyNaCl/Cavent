package repo

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetBriefUser(userId string) (*model.User, error)
	GetUserInterests(userId string) ([]*model.Category, error)
	UpdateUserInterests(userId string, interestIds []string) error
	GetCampusId(userId string) (*uuid.UUID, error)
	UpdateUserCampus(userId string, campusId uuid.UUID) (*model.User, error)
	GetUserProfile(userId string) (*model.User, error)
	UpdateUserProfile(user *model.User) (*model.User, error)
	GetPasswordByUserId(userId string) (string, error)
	UpdateUserPassword(userId, newPassword string) error
	SetUserPassword(userId, newPassword string) error
	HasPassword(userId string) error
}