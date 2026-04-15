package subscription

import (
	"github.com/ekkx/tcm-platform/internal/config"
	"github.com/ekkx/tcm-platform/internal/gen/pb/subscription/v1/subscriptionv1connect"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/modules/subscription/handler"
	"github.com/ekkx/tcm-platform/internal/modules/subscription/repository"
	"github.com/ekkx/tcm-platform/internal/modules/subscription/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
	stripe "github.com/stripe/stripe-go/v85"
)

func InitModule(dbPool *pgxpool.Pool, stripeCfg config.StripeConfig) subscriptionv1connect.SubscriptionServiceHandler {
	querier := sqlc.New(dbPool)
	subRepo := repository.New(querier)
	stripeClient := stripe.NewClient(stripeCfg.SecretKey)
	subUC := usecase.New(subRepo, stripeClient, stripeCfg)
	return handler.New(subUC)
}
