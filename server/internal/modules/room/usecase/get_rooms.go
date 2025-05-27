package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/internal/core/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/core/types"
)

func (uc *Usecase) GetRooms(ctx context.Context) ([]entity.Room, error) {
	tcmRooms := uc.tcmClient.GetRooms()

	// TODO: refactor to a separate function or package if needed
	rooms := make([]entity.Room, len(tcmRooms))
	for i, tcmRoom := range tcmRooms {
		rooms[i] = entity.Room{
			ID:          tcmRoom.ID,
			Name:        tcmRoom.Name,
			PianoType:   convertRoomPianoType(tcmRoom.PianoType),
			PianoNumber: int32(tcmRoom.PianoNumber),
			IsClassroom: tcmRoom.IsClassroom,
			IsBasement:  tcmRoom.IsBasement,
			CampusType:  convertCampusType(tcmRoom.Campus),
			Floor:       int32(tcmRoom.Floor),
		}
	}

	return rooms, nil
}

func convertRoomPianoType(pianoType tcmrsv.RoomPianoType) types.PianoType {
	switch pianoType {
	case tcmrsv.RoomPianoTypeGrand:
		return types.PianoTypeGrand
	case tcmrsv.RoomPianoTypeUpright:
		return types.PianoTypeUpright
	case tcmrsv.RoomPianoTypeUnknown:
		return types.PianoTypeNone
	default:
		return types.PianoTypeNone
	}
}

func convertCampusType(campusType tcmrsv.Campus) types.CampusType {
	switch campusType {
	case tcmrsv.CampusIkebukuro:
		return types.CampusTypeIkebukuro
	case tcmrsv.CampusNakameguro:
		return types.CampusTypeNakameguro
	default:
		return types.CampusTypeUnknown
	}
}
