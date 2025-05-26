package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func (c *CategoryRepository) GetAllCategory() ([]*model.CategoryType, error) {
	var categoryTypes []*model.CategoryType
	err := c.db.
		Preload("Categories", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "category_type_id")
		}).
		Find(&categoryTypes).Error

	if err != nil {
		return nil, err
	}

	return categoryTypes, nil
}

func NewCategoryRepository(db *gorm.DB) repo.CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}
