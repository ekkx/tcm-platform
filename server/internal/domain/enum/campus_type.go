package enum

type CampusType int

const (
	CampusTypeUnknown    CampusType = 0
	CampusTypeNakameguro CampusType = 1
	CampusTypeIkebukuro  CampusType = 2
)

func (ct CampusType) IsValid() bool {
	_, ok := map[CampusType]struct{}{
		CampusTypeUnknown:    {},
		CampusTypeNakameguro: {},
		CampusTypeIkebukuro:  {},
	}[ct]
	return ok
}
