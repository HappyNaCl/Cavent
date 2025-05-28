package command

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type GetEventsCommand struct {
	Limit int
}

type GetEventsCommandResult struct {
	Result []*common.BriefEventResult 
}