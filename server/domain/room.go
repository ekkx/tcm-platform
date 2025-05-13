package domain

type Room struct {
	ID          string
	Name        string
	Campus      Campus
	PianoType   PianoType
	PianoNumber int
	IsClassroom bool
	IsBasement  bool
	Floor       int
}
