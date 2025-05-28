package enum

type PianoType int32

const (
	PianoTypeUnknown PianoType = 0
	PianoTypeGrand   PianoType = 1
	PianoTypeUpright PianoType = 2
	PianoTypeNone    PianoType = 3
)
