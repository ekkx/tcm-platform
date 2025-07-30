package entity

import "github.com/ekkx/tcmrsv-web/internal/domain/enum"

type Room struct {
	ID          string
	Name        string
	PianoType   enum.PianoType
	PianoNumber int
	IsClassroom bool
	IsBasement  bool
	CampusType  enum.CampusType
	Floor       int
}
