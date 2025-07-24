package mapper

import (
	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
)

func ToRoom(room *tcmrsv.Room) *entity.Room {
	if room == nil {
		return nil
	}

	var pianoType enum.PianoType
	switch room.PianoType {
	case tcmrsv.RoomPianoTypeGrand:
		pianoType = enum.PianoTypeGrand
	case tcmrsv.RoomPianoTypeUpright:
		pianoType = enum.PianoTypeUpright
	case tcmrsv.RoomPianoTypeNone:
		pianoType = enum.PianoTypeNone
	default:
		pianoType = enum.PianoTypeUnknown
	}

	var campusType enum.CampusType
	switch room.Campus {
	case tcmrsv.CampusIkebukuro:
		campusType = enum.CampusTypeIkebukuro
	case tcmrsv.CampusNakameguro:
		campusType = enum.CampusTypeNakameguro
	default:
		campusType = enum.CampusTypeUnknown
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
