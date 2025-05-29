package command

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type GetUserInterestsCommand struct {
	UserId string `json:"userId"`
}

type GetUserInterestsCommandResult struct {
	Result []*common.CategoryResult
}