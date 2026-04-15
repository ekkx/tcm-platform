package webhook

import (
	"context"
	"log/slog"

	"github.com/ekkx/tcm-platform/internal/modules/subscription/repository"
	stripe "github.com/stripe/stripe-go/v85"
)

func handleSubscriptionDeleted(ctx context.Context, subRepo repository.Repository, stripeSub *stripe.Subscription) {
	if stripeSub.Customer == nil {
		return
	}

	sub, err := subRepo.GetSubscriptionByStripeCustomerID(ctx, stripeSub.Customer.ID)
	if err != nil || sub == nil {
		return
	}

	if err := subRepo.DeleteSubscription(ctx, sub.ID); err != nil {
		slog.Error("subscription deleted: failed to delete", "error", err)
	}
}
