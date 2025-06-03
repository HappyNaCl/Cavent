package service

import (
	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
	"github.com/HappyNaCl/Cavent/backend/internal/app/mapper"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb"
	"gorm.io/gorm"
)

type CampusService struct {
	campusRepo repo.CampusRepository
}

func NewCampusService(db *gorm.DB) *CampusService {
	return &CampusService{
		campusRepo: postgresdb.NewCampusGormRepo(db),
	}
}

func (s *CampusService) GetAllCampus(com *command.GetAllCampusCommand) (*command.GetAllCampusCommandResult, error) {
	campuses, err := s.campusRepo.GetAllCampus()
	if err != nil {
		return nil, err
	}
	
	result := &command.GetAllCampusCommandResult{
		Result: make([]*common.CampusResult, 0, len(campuses)),
	}
	for _, campus := range campuses {
		result.Result = append(result.Result, mapper.NewCampusResultFromCampus(campus))
	}

	return result, nil
}