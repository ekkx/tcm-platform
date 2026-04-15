package mapper

import (
	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/domain/valueobject"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
)

func ToSubscription(sub *sqlc.Subscription) *entity.Subscription {
	if sub == nil {
		return nil
	}

	var plan valueobject.PlanType
	switch sub.Plan {
	case sqlc.PlanTypeUnlimited:
		plan = valueobject.PlanTypeUnlimited
	case sqlc.PlanTypeLite:
		plan = valueobject.PlanTypeLite
	case sqlc.PlanTypeStandard:
		plan = valueobject.PlanTypeStandard
	case sqlc.PlanTypePro:
		plan = valueobject.PlanTypePro
	default:
		plan = valueobject.PlanTypeLite
	}

	var monthlyHours *int
	if sub.MonthlyHours != nil {
		h := int(*sub.MonthlyHours)
		monthlyHours = &h
	}

	return &entity.Subscription{
		ID:                   sub.ID,
		UserID:               sub.UserID,
		StripeCustomerID:     sub.StripeCustomerID,
		StripeSubscriptionID: sub.StripeSubscriptionID,
		StripePriceID:        sub.StripePriceID,
		Plan:                 plan,
		MonthlyHours:         monthlyHours,
		Status:               sub.Status,
		CurrentPeriodStart:   sub.CurrentPeriodStart,
		CurrentPeriodEnd:     sub.CurrentPeriodEnd,
		CreateTime:           sub.CreateTime,
		UpdateTime:           sub.UpdateTime,
	}
}
