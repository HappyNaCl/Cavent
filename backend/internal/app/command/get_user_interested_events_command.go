package command

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type GetUserInterestedEventsCommand struct {
	UserId string
}

type GetUserInterestedEventsCommandResult struct {
	Result []*common.BriefEventResult 
}