package service

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/mapper"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CategoryService struct {
	categoryRepo repo.CategoryRepository
}

func NewCategoryService(db *gorm.DB, redis *redis.Client) *CategoryService {
	return &CategoryService{
		categoryRepo: postgresdb.NewCategoryGormRepo(db),
	}
}

func (cs *CategoryService) GetAllCategory() (*command.GetCategoriesCommandResult, error) {
	categoryTypes, err := cs.categoryRepo.GetAllCategory()
	if err != nil {
		return nil, err
	}

	return &command.GetCategoriesCommandResult{
		CategoryTypes: mapper.NewCategoryTypeResultFromCategoryTypes(categoryTypes),
	}, nil
}