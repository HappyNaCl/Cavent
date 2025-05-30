package command

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"


type GetRandomEventsCommand struct {}

type GetRandomEventsCommandResult struct {
	Result []*common.BriefEventResult 
}