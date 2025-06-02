package command

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type GetAllEventCommand struct {
	Limit int
	Sort string
	Page int
	Filters []string
	FilterArgs [][]interface{}
	UserId *string
}

type GetAllEventCommandResult struct {
	Limit int
	Page int
	TotalPage int
	TotalRows int
	Result []*common.BriefEventResult
}