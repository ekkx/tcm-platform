package usecase

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/config"
	"github.com/ekkx/tcm-platform/internal/modules/subscription/repository"
	stripe "github.com/stripe/stripe-go/v85"
)

type UseCase interface {
	GetSubscription(ctx context.Context, input *GetSubscriptionInput) (*GetSubscriptionOutput, error)
	GetUsage(ctx context.Context, input *GetUsageInput) (*GetUsageOutput, error)
	CreateCheckoutSession(ctx context.Context, input *CreateCheckoutSessionInput) (*CreateCheckoutSessionOutput, error)
	CreatePortalSession(ctx context.Context, input *CreatePortalSessionInput) (*CreatePortalSessionOutput, error)
}

type UseCaseImpl struct {
	subRepo      repository.Repository
	stripeClient *stripe.Client
	stripeCfg    config.StripeConfig
}

func New(
	subRepo repository.Repository,
	stripeClient *stripe.Client,
	stripeCfg config.StripeConfig,
) UseCase {
	return &UseCaseImpl{
		subRepo:      subRepo,
		stripeClient: stripeClient,
		stripeCfg:    stripeCfg,
	}
}
