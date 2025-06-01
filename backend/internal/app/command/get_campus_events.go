package command

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/google/uuid"
)


type GetCampusEventsCommand struct {
	CampusId uuid.UUID
	UserId *string
	Limit int
	Page int
}

type GetCampusEventsCommandResult struct {
	Result []*common.BriefEventResult
	TotalRows int64
	TotalPages int
	Page int
	Limit int
}

