package mapper

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
)

func NewCampusResultFromCampus(campus *model.Campus) *common.CampusResult {
	if campus == nil {
		return nil
	}

	return &common.CampusResult{
		Id:          campus.Id,
		Name:        campus.Name,
		LogoUrl:     campus.LogoUrl,
		Description: campus.Description,
		InviteCode:  campus.InviteCode,
	}
}