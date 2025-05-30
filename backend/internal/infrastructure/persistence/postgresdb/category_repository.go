package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"gorm.io/gorm"
)

type CategoryGormRepo struct {
	db *gorm.DB
}

func (c *CategoryGormRepo) GetAllCategoryTypes() ([]*model.CategoryType, error) {
	var categoryTypes []*model.CategoryType
	err := c.db.
		Find(&categoryTypes).Error

	if err != nil {
		return nil, err
	}

	return categoryTypes, nil
}

func (c *CategoryGormRepo) GetAllCategory() ([]*model.CategoryType, error) {
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

func NewCategoryGormRepo(db *gorm.DB) repo.CategoryRepository {
	return &CategoryGormRepo{
		db: db,
	}
}
