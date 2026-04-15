package repository

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
)

type Repository interface {
	GetSubscriptionByUserID(ctx context.Context, userID ulid.ULID) (*entity.Subscription, error)
	GetSubscriptionByStripeCustomerID(ctx context.Context, stripeCustomerID string) (*entity.Subscription, error)
	CreateSubscription(ctx context.Context, params *CreateSubscriptionParams) (*ulid.ULID, error)
	UpdateSubscription(ctx context.Context, params *UpdateSubscriptionParams) error
	DeleteSubscription(ctx context.Context, id ulid.ULID) error
	GetUsedMinutesByUserID(ctx context.Context, userID ulid.ULID) (int32, error)
}

type RepositoryImpl struct {
	querier sqlc.Querier
}

func New(querier sqlc.Querier) Repository {
	return &RepositoryImpl{
		querier: querier,
	}
}
