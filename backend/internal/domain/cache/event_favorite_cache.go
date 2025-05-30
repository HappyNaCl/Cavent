package cache

import (
	"context"

	"github.com/google/uuid"
)

type EventFavoriteCache interface {
	IncrementEventFavoriteCount(ctx context.Context, eventID uuid.UUID) (int64, error)
	DecrementEventFavoriteCount(ctx context.Context, eventID uuid.UUID) (int64, error)
	GetEventFavoriteCount(ctx context.Context, eventID uuid.UUID) (int64, error)
}