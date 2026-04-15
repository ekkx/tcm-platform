package presenter

import (
	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/domain/valueobject"
	subscriptionv1 "github.com/ekkx/tcm-platform/internal/gen/pb/subscription/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toPlanType(plan valueobject.PlanType) subscriptionv1.PlanType {
	switch plan {
	case valueobject.PlanTypeUnlimited:
		return subscriptionv1.PlanType_PLAN_TYPE_UNLIMITED
	case valueobject.PlanTypeLite:
		return subscriptionv1.PlanType_PLAN_TYPE_LITE
	case valueobject.PlanTypeStandard:
		return subscriptionv1.PlanType_PLAN_TYPE_STANDARD
	case valueobject.PlanTypePro:
		return subscriptionv1.PlanType_PLAN_TYPE_PRO
	default:
		return subscriptionv1.PlanType_PLAN_TYPE_UNSPECIFIED
	}
}

func ToSubscription(sub *entity.Subscription) *subscriptionv1.Subscription {
	if sub == nil {
		return nil
	}

	out := &subscriptionv1.Subscription{
		Id:         sub.ID.String(),
		UserId:     sub.UserID.String(),
		Plan:       toPlanType(sub.Plan),
		Status:     sub.Status,
		CreateTime: timestamppb.New(sub.CreateTime),
	}

	if sub.MonthlyHours != nil {
		h := int32(*sub.MonthlyHours)
		out.MonthlyHours = &h
	}

	if sub.CurrentPeriodStart != nil {
		out.CurrentPeriodStart = timestamppb.New(*sub.CurrentPeriodStart)
	}

	if sub.CurrentPeriodEnd != nil {
		out.CurrentPeriodEnd = timestamppb.New(*sub.CurrentPeriodEnd)
	}

	return out
}
