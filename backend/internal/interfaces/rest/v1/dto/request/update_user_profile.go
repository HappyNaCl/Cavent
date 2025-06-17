package request

import "github.com/HappyNaCl/Cavent/backend/internal/app/command"

type UpdateUserProfileRequest struct {
	Name        string  `json:"name" form:"name" binding:"required"`
	Description *string `json:"description" form:"description"`
	PhoneNumber *string `json:"phoneNumber" form:"phoneNumber"`
	Address     *string `json:"address" form:"address"`
}

func (r UpdateUserProfileRequest) ToCommand() *command.UpdateUserProfileCommand {
	return &command.UpdateUserProfileCommand{
		Name:        r.Name,
		Description: r.Description,
		PhoneNumber: r.PhoneNumber,
		Address:     r.Address,
	}
}