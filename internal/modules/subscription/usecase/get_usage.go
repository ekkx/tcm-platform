package usecase

import (
	"context"

	"github.com/ekkx/tcm-platform/internal/platform/errs"
)

func (uc *UseCaseImpl) GetUsage(ctx context.Context, input *GetUsageInput) (*GetUsageOutput, error) {
	sub, err := uc.subRepo.GetSubscriptionByUserID(ctx, input.Actor.ID)
	if err != nil {
		return nil, err
	}
	if sub == nil {
		return nil, errs.ErrSubscriptionNotFound
	}

	usedMinutes, err := uc.subRepo.GetUsedMinutesByUserID(ctx, input.Actor.ID)
	if err != nil {
		return nil, err
	}

	var totalMinutes *int32
	if sub.MonthlyHours != nil {
		t := int32(*sub.MonthlyHours * 60)
		totalMinutes = &t
	}

	return &GetUsageOutput{
		UsedMinutes:  usedMinutes,
		TotalMinutes: totalMinutes,
	}, nil
}
