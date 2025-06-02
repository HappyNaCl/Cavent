package command

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type GetEventsCommand struct {
	Limit int
	UserId *string
}

type GetEventsCommandResult struct {
	Result []*common.BriefEventResult 
}