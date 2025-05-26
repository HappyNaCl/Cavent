package command

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
)

type GetBriefUserCommand struct {
	UserId string
}

type GetBriefUserCommandResult struct {
	Result common.BriefUserResult
}