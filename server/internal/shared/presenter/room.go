package presenter

import (
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	roomv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/room/v1"
)

func ToRoom(room *entity.Room) *roomv1.Room {
	if room == nil {
		return nil
	}

	var pianoType roomv1.PianoType
	switch room.PianoType {
	case enum.PianoTypeGrand:
		pianoType = roomv1.PianoType_PIANO_TYPE_GRAND
	case enum.PianoTypeUpright:
		pianoType = roomv1.PianoType_PIANO_TYPE_UPRIGHT
	case enum.PianoTypeNone:
		pianoType = roomv1.PianoType_PIANO_TYPE_NONE
	default:
		pianoType = roomv1.PianoType_PIANO_TYPE_UNSPECIFIED
	}

	var campusType roomv1.CampusType
	switch room.CampusType {
	case enum.CampusTypeIkebukuro:
		campusType = roomv1.CampusType_CAMPUS_TYPE_IKEBUKURO
	case enum.CampusTypeNakameguro:
		campusType = roomv1.CampusType_CAMPUS_TYPE_NAKAMEGURO
	default:
		campusType = roomv1.CampusType_CAMPUS_TYPE_UNSPECIFIED
	}

	return &roomv1.Room{
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

func ToRoomList(rooms []*entity.Room) []*roomv1.Room {
	if rooms == nil {
		return nil
	}
	result := make([]*roomv1.Room, len(rooms))
	for i, room := range rooms {
		result[i] = ToRoom(room)
	}
	return result
}
