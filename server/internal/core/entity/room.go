package entity

import "github.com/ekkx/tcmrsv-web/server/internal/core/types"

type Room struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	PianoType   types.PianoType  `json:"piano_type"`
	PianoNumber int32            `json:"piano_number"`
	IsClassroom bool             `json:"is_classroom"`
	IsBasement  bool             `json:"is_basement"`
	CampusType  types.CampusType `json:"campus_type"`
	Floor       int32            `json:"floor"`
}
