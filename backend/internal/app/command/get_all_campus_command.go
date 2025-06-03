package command

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type GetAllCampusCommand struct{}

type GetAllCampusCommandResult struct {
	Result []*common.CampusResult `json:"campus"`
}