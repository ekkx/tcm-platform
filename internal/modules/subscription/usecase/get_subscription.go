package usecase

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/platform/errs"
)

func (uc *UseCaseImpl) GetSubscription(ctx context.Context, input *GetSubscriptionInput) (*GetSubscriptionOutput, error) {
	sub, err := uc.subRepo.GetSubscriptionByUserID(ctx, input.Actor.ID)
	if err != nil {
		return nil, err
	}
	if sub == nil {
		return nil, errs.ErrSubscriptionNotFound
	}
	return &GetSubscriptionOutput{Subscription: sub}, nil
}
