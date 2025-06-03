package command

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type UpdateUserCampusCommand struct {
	InviteCode string `json:"inviteCode" form:"inviteCode" binding:"required"`
	UserId string `json:"userId" form:"userId"`
}

type UpdateUserCampusCommandResult struct {
	User *common.BriefUserResult
}