package entity

import "github.com/ekkx/tcmrsv-web/server/internal/domain/enum"

type Room struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	PianoType   enum.PianoType  `json:"piano_type"`
	PianoNumber int             `json:"piano_number"`
	IsClassroom bool            `json:"is_classroom"`
	IsBasement  bool            `json:"is_basement"`
	CampusType  enum.CampusType `json:"campus_type"`
	Floor       int             `json:"floor"`
}
