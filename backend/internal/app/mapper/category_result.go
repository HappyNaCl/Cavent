package mapper

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
)

func NewCategoryTypeResultFromCategoryTypes(categoryTypes []*model.CategoryType) []*common.CategoryTypeResult {
	if categoryTypes == nil {
		return nil
	}

	categoryTypeResults := make([]*common.CategoryTypeResult, 0, len(categoryTypes))
	for _, categoryType := range categoryTypes {
		categories := make([]*common.CategoryResult, 0, len(categoryType.Categories))
		for _, category := range categoryType.Categories {
			categories = append(categories, &common.CategoryResult{
				Id:   category.Id,
				Name: category.Name,
			})
		}

		categoryTypeResults = append(categoryTypeResults, &common.CategoryTypeResult{
			Id:       categoryType.Id.String(),
			Name:     categoryType.Name,
			Categories: categories,
		})
	}

	return categoryTypeResults
}

func NewCategoriesResultFromCategory(category []*model.Category) []*common.CategoryResult {
	if category == nil {
		return nil
	}

	categoryResults := make([]*common.CategoryResult, 0, len(category))
	for _, cat := range category {
		categoryResults = append(categoryResults, &common.CategoryResult{
			Id:   cat.Id,
			Name: cat.Name,
		})
	}

	return categoryResults
}