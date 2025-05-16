package repo

import "github.com/HappyNaCl/Cavent/backend/internal/domain/model"

type UserRepository interface {
	FindByEmail(email string) (*model.User, error)
}