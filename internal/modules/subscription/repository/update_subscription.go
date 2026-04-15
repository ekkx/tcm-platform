package repository

import (
	"context"
	"time"

	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
)

type UpdateSubscriptionParams struct {
	ID                   ulid.ULID
	StripeSubscriptionID *string
	StripePriceID        *string
	Plan                 sqlc.NullPlanType
	MonthlyHours         *int32
	Status               *string
	CurrentPeriodStart   *time.Time
	CurrentPeriodEnd     *time.Time
}

func (repo *RepositoryImpl) UpdateSubscription(ctx context.Context, params *UpdateSubscriptionParams) error {
	_, err := repo.querier.UpdateSubscription(ctx, sqlc.UpdateSubscriptionParams{
		ID:                   params.ID,
		StripeSubscriptionID: params.StripeSubscriptionID,
		StripePriceID:        params.StripePriceID,
		Plan:                 params.Plan,
		MonthlyHours:         params.MonthlyHours,
		Status:               params.Status,
		CurrentPeriodStart:   params.CurrentPeriodStart,
		CurrentPeriodEnd:     params.CurrentPeriodEnd,
	})
	return err
}
