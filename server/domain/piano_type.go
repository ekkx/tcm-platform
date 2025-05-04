package domain

type PianoType string

const (
	PianoTypeGrand   PianoType = "grand"
	PianoTypeUpright PianoType = "upright"
	PianoTypeNone    PianoType = "none"
)

func (r PianoType) IsValid() bool {
	switch r {
	case PianoTypeGrand, PianoTypeUpright, PianoTypeNone:
		return true
	default:
		return false
	}
}
