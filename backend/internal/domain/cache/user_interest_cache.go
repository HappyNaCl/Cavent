package cache

import (
	"context"

	"github.com/HappyNaCl/Cavent/backend/internal/app/common"
)

type UserInterestCache interface {
	GetUserInterest(ctx context.Context, userID string) ([]*common.CategoryResult, error)
	SetUserInterest(ctx context.Context, userID string, categories []*common.CategoryResult) error
	Invalidate(ctx context.Context, userID string) error
}