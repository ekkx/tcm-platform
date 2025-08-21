package presenter

import (
	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/domain/valueobject"
	roomv1 "github.com/ekkx/tcm-platform/internal/gen/pb/room/v1"
)

func ToRoom(room *entity.Room) *roomv1.Room {
	if room == nil {
		return nil
	}

	var pianoType roomv1.PianoType
	switch room.PianoType {
	case valueobject.PianoTypeGrand:
		pianoType = roomv1.PianoType_PIANO_TYPE_GRAND
	case valueobject.PianoTypeUpright:
		pianoType = roomv1.PianoType_PIANO_TYPE_UPRIGHT
	case valueobject.PianoTypeNone:
		pianoType = roomv1.PianoType_PIANO_TYPE_NONE
	default:
		pianoType = roomv1.PianoType_PIANO_TYPE_UNSPECIFIED
	}

	var campusType roomv1.CampusType
	switch room.CampusType {
	case valueobject.CampusTypeIkebukuro:
		campusType = roomv1.CampusType_CAMPUS_TYPE_IKEBUKURO
	case valueobject.CampusTypeNakameguro:
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
