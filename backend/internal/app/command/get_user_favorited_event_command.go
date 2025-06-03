package command

import "github.com/HappyNaCl/Cavent/backend/internal/app/common"

type GetUserFavoritedEventCommand struct {
	UserId string
}

type GetUserFavoritedEventResult struct {
	Result []*common.BriefEventResult
}

