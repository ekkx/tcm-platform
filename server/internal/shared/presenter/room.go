package presenter

import (
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	reservationv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1"
)

func ToRoom(room *entity.Room) *reservationv1.Room {
	if room == nil {
		return nil
	}

	var pianoType reservationv1.PianoType
	switch room.PianoType {
	case enum.PianoTypeGrand:
		pianoType = reservationv1.PianoType_PIANO_TYPE_GRAND
	case enum.PianoTypeUpright:
		pianoType = reservationv1.PianoType_PIANO_TYPE_UPRIGHT
	case enum.PianoTypeNone:
		pianoType = reservationv1.PianoType_PIANO_TYPE_NONE
	default:
		pianoType = reservationv1.PianoType_PIANO_TYPE_UNSPECIFIED
	}

	var campusType reservationv1.CampusType
	switch room.CampusType {
	case enum.CampusTypeIkebukuro:
		campusType = reservationv1.CampusType_CAMPUS_TYPE_IKEBUKURO
	case enum.CampusTypeNakameguro:
		campusType = reservationv1.CampusType_CAMPUS_TYPE_NAKAMEGURO
	default:
		campusType = reservationv1.CampusType_CAMPUS_TYPE_UNSPECIFIED
	}

	return &reservationv1.Room{
		Id:          room.ID,
		Name:        room.Name,
		PianoType:   pianoType,
		PianoCount:  int32(room.PianoNumber),
		IsClassroom: room.IsClassroom,
		IsBasement:  room.IsBasement,
		CampusType:  campusType,
		Floor:       int32(room.Floor),
	}
}

func ToRoomList(rooms []*entity.Room) []*reservationv1.Room {
	if rooms == nil {
		return nil
	}
	result := make([]*reservationv1.Room, len(rooms))
	for i, room := range rooms {
		result[i] = ToRoom(room)
	}
	return result
}
