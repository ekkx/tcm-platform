package webhook

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/ekkx/tcm-platform/internal/config"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/modules/subscription/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	stripe "github.com/stripe/stripe-go/v85"
	stripewebhook "github.com/stripe/stripe-go/v85/webhook"
)

func NewHandler(dbPool *pgxpool.Pool, stripeCfg config.StripeConfig, stripeClient *stripe.Client) http.HandlerFunc {
	querier := sqlc.New(dbPool)
	subRepo := repository.New(querier)

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", http.StatusBadRequest)
			return
		}

		event, err := stripewebhook.ConstructEvent(body, r.Header.Get("Stripe-Signature"), stripeCfg.WebhookSecret)
		if err != nil {
			slog.Error("webhook signature verification failed", "error", err)
			http.Error(w, "invalid signature", http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		switch event.Type {
		case "checkout.session.completed":
			var session stripe.CheckoutSession
			if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
				slog.Error("failed to unmarshal checkout session", "error", err)
				http.Error(w, "bad payload", http.StatusBadRequest)
				return
			}
			handleCheckoutCompleted(ctx, subRepo, stripeClient, &session, stripeCfg)

		case "customer.subscription.updated":
			var sub stripe.Subscription
			if err := json.Unmarshal(event.Data.Raw, &sub); err != nil {
				slog.Error("failed to unmarshal subscription", "error", err)
				http.Error(w, "bad payload", http.StatusBadRequest)
				return
			}
			handleSubscriptionUpdated(ctx, subRepo, &sub, stripeCfg)

		case "customer.subscription.deleted":
			var sub stripe.Subscription
			if err := json.Unmarshal(event.Data.Raw, &sub); err != nil {
				slog.Error("failed to unmarshal subscription", "error", err)
				http.Error(w, "bad payload", http.StatusBadRequest)
				return
			}
			handleSubscriptionDeleted(ctx, subRepo, &sub)

		case "invoice.paid":
			var inv stripe.Invoice
			if err := json.Unmarshal(event.Data.Raw, &inv); err != nil {
				slog.Error("failed to unmarshal invoice", "error", err)
				http.Error(w, "bad payload", http.StatusBadRequest)
				return
			}
			handleInvoicePaid(ctx, subRepo, &inv)
		}

		w.WriteHeader(http.StatusOK)
	}
}
