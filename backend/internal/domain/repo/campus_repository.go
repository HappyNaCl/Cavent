package repo

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/google/uuid"
)

type CampusRepository interface {
	FindCampusByInviteCode(inviteCode string) (*model.Campus, error)
	GetAllCampus() ([]*model.Campus, error)
	GetCampusById(campusId uuid.UUID) (*model.Campus, error)
}