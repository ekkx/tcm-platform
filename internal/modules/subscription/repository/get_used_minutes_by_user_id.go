package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/platform/ulid"
)

func (repo *RepositoryImpl) GetUsedMinutesByUserID(ctx context.Context, userID ulid.ULID) (int32, error) {
	return repo.querier.GetUsedMinutesByUserID(ctx, userID)
}
