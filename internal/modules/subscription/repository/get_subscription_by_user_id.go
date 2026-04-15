package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/modules/subscription/mapper"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
	"github.com/jackc/pgx/v5"
)

func (repo *RepositoryImpl) GetSubscriptionByUserID(ctx context.Context, userID ulid.ULID) (*entity.Subscription, error) {
	row, err := repo.querier.GetSubscriptionByUserID(ctx, userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return mapper.ToSubscription(&row), nil
}
