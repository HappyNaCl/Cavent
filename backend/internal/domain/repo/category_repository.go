package repo

import "github.com/HappyNaCl/Cavent/backend/internal/domain/model"

type CategoryRepository interface {
	GetAllCategory() ([]*model.CategoryType, error) 
}