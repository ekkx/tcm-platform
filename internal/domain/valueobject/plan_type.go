package valueobject

type PlanType string

const (
	PlanTypeUnlimited PlanType = "unlimited"
	PlanTypeLite      PlanType = "lite"
	PlanTypeStandard  PlanType = "standard"
	PlanTypePro       PlanType = "pro"
)

func (p PlanType) MonthlyHours() *int {
	var hours int
	switch p {
	case PlanTypeUnlimited:
		return nil
	case PlanTypeLite:
		hours = 30
	case PlanTypeStandard:
		hours = 60
	case PlanTypePro:
		hours = 90
	default:
		hours = 30
	}
	return &hours
}
