package mapper

import (
	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/domain/valueobject"
	"github.com/ekkx/tcmrsv"
)

func ToRoom(room *tcmrsv.Room) *entity.Room {
	if room == nil {
		return nil
	}

	var pianoType valueobject.PianoType
	switch room.PianoType {
	case tcmrsv.RoomPianoTypeGrand:
		pianoType = valueobject.PianoTypeGrand
	case tcmrsv.RoomPianoTypeUpright:
		pianoType = valueobject.PianoTypeUpright
	case tcmrsv.RoomPianoTypeNone:
		pianoType = valueobject.PianoTypeNone
	default:
		pianoType = valueobject.PianoTypeUnknown
	}

	var campusType valueobject.CampusType
	switch room.Campus {
	case tcmrsv.CampusIkebukuro:
		campusType = valueobject.CampusTypeIkebukuro
	case tcmrsv.CampusNakameguro:
		campusType = valueobject.CampusTypeNakameguro
	default:
		campusType = valueobject.CampusTypeUnknown
	}

	return &entity.Room{
		ID:          room.ID,
		Name:        room.Name,
		PianoType:   pianoType,
		PianoNumber: room.PianoNumber,
		IsClassroom: room.IsClassroom,
		IsBasement:  room.IsBasement,
		CampusType:  campusType,
		Floor:       room.Floor,
	}
}

func ToRoomList(rooms []*tcmrsv.Room) []*entity.Room {
	if rooms == nil {
		return nil
	}
	result := make([]*entity.Room, len(rooms))
	for i, room := range rooms {
		result[i] = ToRoom(room)
	}
	return result
}
