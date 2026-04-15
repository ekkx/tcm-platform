package webhook

import (
	"context"
	"log/slog"
	"time"

	"github.com/ekkx/tcm-platform/internal/config"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/modules/subscription/repository"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
	stripe "github.com/stripe/stripe-go/v85"
)

func handleCheckoutCompleted(ctx context.Context, subRepo repository.Repository, stripeClient *stripe.Client, session *stripe.CheckoutSession, stripeCfg config.StripeConfig) {
	if session.Customer == nil {
		slog.Error("checkout completed: no customer in session")
		return
	}

	customerID := session.Customer.ID

	// Stripe API から Customer を取得してメタデータを読む
	customer, err := stripeClient.V1Customers.Retrieve(ctx, customerID, &stripe.CustomerRetrieveParams{})
	if err != nil {
		slog.Error("checkout completed: failed to retrieve customer", "customer_id", customerID, "error", err)
		return
	}

	userIDStr, ok := customer.Metadata["tcm_user_id"]
	if !ok || userIDStr == "" {
		slog.Error("checkout completed: no tcm_user_id in customer metadata", "customer_id", customerID)
		return
	}

	userID, err := ulid.Parse(userIDStr)
	if err != nil {
		slog.Error("checkout completed: invalid tcm_user_id", "user_id", userIDStr, "error", err)
		return
	}

	// Subscription ID を取得
	stripeSubID := ""
	if session.Subscription != nil {
		stripeSubID = session.Subscription.ID
	}

	// Stripe API から Subscription を取得して price ID と period を読む
	var priceID string
	var periodStart, periodEnd time.Time
	if stripeSubID != "" {
		stripeSub, err := stripeClient.V1Subscriptions.Retrieve(ctx, stripeSubID, &stripe.SubscriptionRetrieveParams{})
		if err == nil {
			if len(stripeSub.Items.Data) > 0 && stripeSub.Items.Data[0].Price != nil {
				priceID = stripeSub.Items.Data[0].Price.ID
			}
			if len(stripeSub.Items.Data) > 0 {
				periodStart = time.Unix(stripeSub.Items.Data[0].CurrentPeriodStart, 0)
				periodEnd = time.Unix(stripeSub.Items.Data[0].CurrentPeriodEnd, 0)
			}
		}
	}

	plan, monthlyHours := resolvePlan(priceID, stripeCfg)

	// 既存レコードがあるか確認
	sub, err := subRepo.GetSubscriptionByStripeCustomerID(ctx, customerID)
	if err != nil {
		slog.Error("checkout completed: failed to get subscription", "customer_id", customerID, "error", err)
		return
	}

	if sub != nil {
		// 既存レコードを更新
		status := "active"
		if err := subRepo.UpdateSubscription(ctx, &repository.UpdateSubscriptionParams{
			ID:                   sub.ID,
			StripeSubscriptionID: &stripeSubID,
			StripePriceID:        &priceID,
			Plan:                 sqlc.NullPlanType{PlanType: plan, Valid: true},
			MonthlyHours:         int32Ptr(monthlyHours),
			Status:               &status,
			CurrentPeriodStart:   &periodStart,
			CurrentPeriodEnd:     &periodEnd,
		}); err != nil {
			slog.Error("checkout completed: failed to update subscription", "error", err)
		}
		return
	}

	// 新規作成
	if _, err := subRepo.CreateSubscription(ctx, &repository.CreateSubscriptionParams{
		UserID:               userID,
		StripeCustomerID:     &customerID,
		StripeSubscriptionID: &stripeSubID,
		StripePriceID:        &priceID,
		Plan:                 plan,
		MonthlyHours:         int32Ptr(monthlyHours),
		Status:               "active",
		CurrentPeriodStart:   &periodStart,
		CurrentPeriodEnd:     &periodEnd,
	}); err != nil {
		slog.Error("checkout completed: failed to create subscription", "error", err)
	}
}

func int32Ptr(v int32) *int32 {
	return &v
}
