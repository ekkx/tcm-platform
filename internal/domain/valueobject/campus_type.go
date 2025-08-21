package valueobject

type CampusType int

const (
	CampusTypeUnknown CampusType = iota
	CampusTypeNakameguro
	CampusTypeIkebukuro
)

func (ct CampusType) IsValid() bool {
	_, ok := map[CampusType]struct{}{
		CampusTypeNakameguro: {},
		CampusTypeIkebukuro:  {},
	}[ct]
	return ok
}
