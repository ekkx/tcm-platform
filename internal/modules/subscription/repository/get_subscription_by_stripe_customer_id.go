package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/modules/subscription/mapper"
	"github.com/jackc/pgx/v5"
)

func (repo *RepositoryImpl) GetSubscriptionByStripeCustomerID(ctx context.Context, stripeCustomerID string) (*entity.Subscription, error) {
	row, err := repo.querier.GetSubscriptionByStripeCustomerID(ctx, stripeCustomerID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return mapper.ToSubscription(&row), nil
}
