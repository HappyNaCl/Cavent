package request

import "github.com/HappyNaCl/Cavent/backend/internal/app/command"

type UpdateUserInterestRequest struct {
	UserId      string   `json:"userId" form:"userId"`
	CategoryIds []string `json:"categoryIds" form:"categoryIds" binding:"required"`
}

func (r *UpdateUserInterestRequest) ToCommand() *command.UpdateUserInterestCommand {
	return &command.UpdateUserInterestCommand{
		UserId:      r.UserId,
		CategoryIds: r.CategoryIds,
	}
}
