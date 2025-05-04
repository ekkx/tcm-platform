package domain

type Campus string

const (
	CampusIkebukuro  Campus = "1"
	CampusNakameguro Campus = "2"
)

func (c Campus) IsValid() bool {
	switch c {
	case CampusIkebukuro, CampusNakameguro:
		return true
	default:
		return false
	}
}
