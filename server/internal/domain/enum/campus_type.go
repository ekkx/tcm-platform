package enum

type CampusType int32

const (
	CampusTypeUnknown    CampusType = 0
	CampusTypeNakameguro CampusType = 1
	CampusTypeIkebukuro  CampusType = 2
)

func (mu CampusType) IsValid() bool {
	_, ok := map[CampusType]struct{}{
		CampusTypeUnknown:    {},
		CampusTypeNakameguro: {},
		CampusTypeIkebukuro:  {},
	}[mu]
	return ok
}
