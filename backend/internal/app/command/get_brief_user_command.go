package command

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type GetBriefUserCommand struct {
	UserID string
}

type GetBriefUserCommandResult struct {
	Result *common.BriefUserResult
}