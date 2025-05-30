package cache

import (
	"context"

	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
)

type CategoriesCache interface {
	GetAllCategories(ctx context.Context) ([]*common.CategoryResult, error)
	SetAllCategories(ctx context.Context, categories []*common.CategoryResult) error
	Invalidate(ctx context.Context) error
}
