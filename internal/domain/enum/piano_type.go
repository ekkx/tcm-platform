package enum

type PianoType int

const (
	PianoTypeUnknown PianoType = iota
	PianoTypeGrand
	PianoTypeUpright
	PianoTypeNone
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
