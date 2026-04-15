package usecase

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/modules/subscription/repository"
	"github.com/ekkx/tcm-platform/internal/platform/errs"
	stripe "github.com/stripe/stripe-go/v85"
)

func (uc *UseCaseImpl) CreateCheckoutSession(ctx context.Context, input *CreateCheckoutSessionInput) (*CreateCheckoutSessionOutput, error) {
	// 既存サブスクリプションを確認
	sub, err := uc.subRepo.GetSubscriptionByUserID(ctx, input.Actor.ID)
	if err != nil {
		return nil, err
	}

	// unlimited ユーザーは Checkout 不可
	if sub != nil && sub.IsUnlimited() {
		return nil, errs.ErrUnlimitedUserCannotCheckout
	}

	// Stripe サブスクリプションが存在する場合、Stripe 側の実際の状態を確認
	if sub != nil && sub.StripeSubscriptionID != nil {
		stripeSub, err := uc.stripeClient.V1Subscriptions.Retrieve(ctx, *sub.StripeSubscriptionID, &stripe.SubscriptionRetrieveParams{})
		if err == nil && stripeSub.Status == "active" {
			// Stripe 側でアクティブ → プラン変更は Customer Portal で行う
			return nil, errs.ErrAlreadySubscribed
		}
		// キャンセル済み/期限切れなど → DB をリセットして新規 Checkout へ進む
		canceled := "canceled"
		_ = uc.subRepo.UpdateSubscription(ctx, &repository.UpdateSubscriptionParams{
			ID:     sub.ID,
			Status: &canceled,
		})
	}

	var stripeCustomerID string

	if sub != nil && sub.StripeCustomerID != nil {
		stripeCustomerID = *sub.StripeCustomerID
	} else {
		// 新規 Stripe Customer を作成
		customer, err := uc.stripeClient.V1Customers.Create(ctx, &stripe.CustomerCreateParams{
			Metadata: map[string]string{
				"tcm_user_id": input.Actor.ID.String(),
			},
		})
		if err != nil {
			return nil, errs.ErrCheckoutSessionFailed.WithCause(err)
		}
		stripeCustomerID = customer.ID
	}

	// Stripe Checkout Session を作成（新規契約のみ）
	session, err := uc.stripeClient.V1CheckoutSessions.Create(ctx, &stripe.CheckoutSessionCreateParams{
		Customer: stripe.String(stripeCustomerID),
		Mode:     stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		Locale:   stripe.String("ja"),
		LineItems: []*stripe.CheckoutSessionCreateLineItemParams{
			{
				Price:    stripe.String(input.PriceID),
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL: stripe.String(uc.stripeCfg.SuccessURL()),
		CancelURL:  stripe.String(uc.stripeCfg.CancelURL()),
	})
	if err != nil {
		return nil, errs.ErrCheckoutSessionFailed.WithCause(err)
	}

	return &CreateCheckoutSessionOutput{CheckoutURL: session.URL}, nil
}
