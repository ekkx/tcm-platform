package webhook

import (
	"context"
	"log/slog"
	"time"

	"github.com/ekkx/tcm-platform/internal/config"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/modules/subscription/repository"
	stripe "github.com/stripe/stripe-go/v85"
)

func handleSubscriptionUpdated(ctx context.Context, subRepo repository.Repository, stripeSub *stripe.Subscription, stripeCfg config.StripeConfig) {
	if stripeSub.Customer == nil {
		return
	}

	sub, err := subRepo.GetSubscriptionByStripeCustomerID(ctx, stripeSub.Customer.ID)
	if err != nil || sub == nil {
		slog.Error("subscription updated: not found", "customer_id", stripeSub.Customer.ID, "error", err)
		return
	}

	var priceID string
	if len(stripeSub.Items.Data) > 0 && stripeSub.Items.Data[0].Price != nil {
		priceID = stripeSub.Items.Data[0].Price.ID
	}

	plan, monthlyHours := resolvePlan(priceID, stripeCfg)
	status := string(stripeSub.Status)

	var periodStart, periodEnd time.Time
	if len(stripeSub.Items.Data) > 0 {
		periodStart = time.Unix(stripeSub.Items.Data[0].CurrentPeriodStart, 0)
		periodEnd = time.Unix(stripeSub.Items.Data[0].CurrentPeriodEnd, 0)
	}

	if err := subRepo.UpdateSubscription(ctx, &repository.UpdateSubscriptionParams{
		ID:                 sub.ID,
		StripePriceID:      &priceID,
		Plan:               sqlc.NullPlanType{PlanType: plan, Valid: true},
		MonthlyHours:       &monthlyHours,
		Status:             &status,
		CurrentPeriodStart: &periodStart,
		CurrentPeriodEnd:   &periodEnd,
	}); err != nil {
		slog.Error("subscription updated: failed to update", "error", err)
	}
}
