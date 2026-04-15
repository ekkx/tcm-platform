package webhook

import (
	"github.com/ekkx/tcm-platform/internal/config"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
)

func resolvePlan(priceID string, cfg config.StripeConfig) (sqlc.PlanType, int32) {
	switch priceID {
	case cfg.PriceLite:
		return sqlc.PlanTypeLite, 30
	case cfg.PriceStandard:
		return sqlc.PlanTypeStandard, 60
	case cfg.PricePro:
		return sqlc.PlanTypePro, 90
	default:
		return sqlc.PlanTypeLite, 30
	}
}
