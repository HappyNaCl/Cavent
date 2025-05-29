package command

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type GetCategoryTypesCommand struct {}

type GetCategoryTypesCommandResult struct {
	Result []*common.CategoryTypeResult	
}