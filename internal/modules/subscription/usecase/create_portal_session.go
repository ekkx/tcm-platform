package usecase

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/platform/errs"
	stripe "github.com/stripe/stripe-go/v85"
)

func (uc *UseCaseImpl) CreatePortalSession(ctx context.Context, input *CreatePortalSessionInput) (*CreatePortalSessionOutput, error) {
	sub, err := uc.subRepo.GetSubscriptionByUserID(ctx, input.Actor.ID)
	if err != nil {
		return nil, err
	}
	if sub == nil {
		return nil, errs.ErrSubscriptionNotFound
	}
	if sub.IsUnlimited() {
		return nil, errs.ErrUnlimitedUserCannotCheckout
	}
	if sub.StripeCustomerID == nil {
		return nil, errs.ErrSubscriptionNotFound.WithMessage("no stripe customer associated")
	}

	session, err := uc.stripeClient.V1BillingPortalSessions.Create(ctx, &stripe.BillingPortalSessionCreateParams{
		Customer:  sub.StripeCustomerID,
		Locale:    stripe.String("ja"),
		ReturnURL: stripe.String(uc.stripeCfg.PortalReturnURL()),
	})
	if err != nil {
		return nil, errs.ErrPortalSessionFailed.WithCause(err)
	}

	return &CreatePortalSessionOutput{PortalURL: session.URL}, nil
}
