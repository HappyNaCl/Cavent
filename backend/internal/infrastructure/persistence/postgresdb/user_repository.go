package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"gorm.io/gorm"
)

type UserGormRepo struct {
	db    *gorm.DB
}

// FindByEmail implements repo.UserRepository.
func (u *UserGormRepo) FindByEmail(email string) (*model.User, error) {
	panic("unimplemented")
}

func NewUserGormRepo(db *gorm.DB) repo.UserRepository {
	return &UserGormRepo{
		db:    db,
	}
}
