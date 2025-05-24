package repo

import "github.com/HappyNaCl/Cavent/backend/internal/domain/model"

type CampusRepository interface {
	FindCampusByInviteCode(inviteCode string) (*model.Campus, error)
}