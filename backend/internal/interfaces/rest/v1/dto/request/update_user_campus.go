package request

import "github.com/HappyNaCl/Cavent/backend/internal/app/command"

type UpdateUserCampusRequest struct {
	InviteCode string `json:"inviteCode" form:"inviteCode" binding:"required"`
	UserId     string `json:"userId"`
}

func (r UpdateUserCampusRequest) ToCommand() *command.UpdateUserCampusCommand {
	return &command.UpdateUserCampusCommand{
		InviteCode: r.InviteCode,
		UserId:     r.UserId,
	}
}