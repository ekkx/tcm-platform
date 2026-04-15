package entity

import (
	"time"

	"github.com/ekkx/tcm-platform/internal/domain/valueobject"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
)

type Subscription struct {
	ID                   ulid.ULID
	UserID               ulid.ULID
	StripeCustomerID     *string
	StripeSubscriptionID *string
	StripePriceID        *string
	Plan                 valueobject.PlanType
	MonthlyHours         *int
	Status               string
	CurrentPeriodStart   *time.Time
	CurrentPeriodEnd     *time.Time
	CreateTime           time.Time
	UpdateTime           time.Time
}

func (s *Subscription) IsUnlimited() bool {
	return s.Plan == valueobject.PlanTypeUnlimited
}

func (s *Subscription) IsActive() bool {
	return s.Status == "active"
}
