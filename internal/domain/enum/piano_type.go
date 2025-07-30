package enum

type PianoType int

const (
	PianoTypeUnknown PianoType = 0
	PianoTypeGrand   PianoType = 1
	PianoTypeUpright PianoType = 2
	PianoTypeNone    PianoType = 3
)

func (pt PianoType) IsValid() bool {
	_, ok := map[PianoType]struct{}{
		PianoTypeUnknown: {},
		PianoTypeGrand:   {},
		PianoTypeUpright: {},
		PianoTypeNone:    {},
	}[pt]
	return ok
}
