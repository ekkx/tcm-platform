package webhook

import (
	"context"
	"log/slog"

	"github.com/ekkx/tcm-platform/internal/modules/subscription/repository"
	stripe "github.com/stripe/stripe-go/v85"
)

func handleInvoicePaid(ctx context.Context, subRepo repository.Repository, inv *stripe.Invoice) {
	if inv.Customer == nil {
		return
	}

	sub, err := subRepo.GetSubscriptionByStripeCustomerID(ctx, inv.Customer.ID)
	if err != nil || sub == nil {
		return
	}

	status := "active"
	if err := subRepo.UpdateSubscription(ctx, &repository.UpdateSubscriptionParams{
		ID:     sub.ID,
		Status: &status,
	}); err != nil {
		slog.Error("invoice paid: failed to update", "error", err)
	}
}
