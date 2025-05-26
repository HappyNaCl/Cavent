package command

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
)

type GetCategoriesCommand struct{}

type GetCategoriesCommandResult struct {
	CategoryTypes []*common.CategoryTypeResult
}