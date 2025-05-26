package mapper

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
)

func NewCategoryResultFromCategoryTypes(categoryTypes []*model.CategoryType) []*common.CategoryTypeResult {
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