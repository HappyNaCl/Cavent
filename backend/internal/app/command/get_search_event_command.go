package command

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type GetSearchEventCommand struct {
	Query string `json:"query"`
}

type GetSearchEventCommandResult struct {
	Result []*common.EventSearchResult `json:"events"`
}