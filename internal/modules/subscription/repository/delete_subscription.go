package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/platform/ulid"
)

func (repo *RepositoryImpl) DeleteSubscription(ctx context.Context, id ulid.ULID) error {
	return repo.querier.DeleteSubscription(ctx, id)
}
