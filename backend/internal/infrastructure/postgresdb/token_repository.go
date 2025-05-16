package postgresdb

import (
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TokenRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (t *TokenRepository) CheckToken(token string) (bool, error) {
	var tokenModel *model.RefreshToken

	err := t.db.Where("token = ?", token).First(&tokenModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		t.logger.Sugar().Errorf("Error checking token: %v", err)
		return false, err
	}

	if tokenModel.Token != token {
		return false, nil
	}

	return true, nil
}

// IssueToken implements repo.TokenRepository.
func (t *TokenRepository) IssueToken(token *model.RefreshToken) (*string, error) {
	err := t.db.Create(token).Error
	if err != nil {
		t.logger.Sugar().Errorf("Error creating token: %v", err)
		return nil, err	
	}

	return &token.Token, nil
}

func (t *TokenRepository) RevokeToken(token string) error {
	panic("unimplemented")
}

func NewTokenRepository(db *gorm.DB, logger *zap.Logger) repo.TokenRepository {
	return &TokenRepository{
		db:     db,
		logger: logger,
	}
}
