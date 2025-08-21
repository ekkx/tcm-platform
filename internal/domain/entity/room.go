package entity

import "github.com/ekkx/tcm-platform/internal/domain/valueobject"

type Room struct {
	ID          string
	Name        string
	PianoType   valueobject.PianoType
	PianoNumber int
	IsClassroom bool
	IsBasement  bool
	CampusType  valueobject.CampusType
	Floor       int
}
