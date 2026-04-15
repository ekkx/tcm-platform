package repository

import (
	"context"
	"time"

	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
)

type CreateSubscriptionParams struct {
	ID                   ulid.ULID
	UserID               ulid.ULID
	StripeCustomerID     *string
	StripeSubscriptionID *string
	StripePriceID        *string
	Plan                 sqlc.PlanType
	MonthlyHours         *int32
	Status               string
	CurrentPeriodStart   *time.Time
	CurrentPeriodEnd     *time.Time
}

func (repo *RepositoryImpl) CreateSubscription(ctx context.Context, params *CreateSubscriptionParams) (*ulid.ULID, error) {
	if params.ID.IsZero() {
		params.ID = ulid.New()
	}

	id, err := repo.querier.CreateSubscription(ctx, sqlc.CreateSubscriptionParams{
		ID:                   params.ID,
		UserID:               params.UserID,
		StripeCustomerID:     params.StripeCustomerID,
		StripeSubscriptionID: params.StripeSubscriptionID,
		StripePriceID:        params.StripePriceID,
		Plan:                 params.Plan,
		MonthlyHours:         params.MonthlyHours,
		Status:               params.Status,
		CurrentPeriodStart:   params.CurrentPeriodStart,
		CurrentPeriodEnd:     params.CurrentPeriodEnd,
	})
	if err != nil {
		return nil, err
	}

	return &id, nil
}
