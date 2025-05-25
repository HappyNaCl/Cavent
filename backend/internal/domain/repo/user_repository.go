package repo

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetBriefUser(userId uuid.UUID) (*model.User, error)
}