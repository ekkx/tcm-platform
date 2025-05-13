package room

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/domain"
)

type GetRoomsOutput struct {
	Rooms []domain.Room
}

func (uc *RoomUsecaseImpl) GetRooms(ctx context.Context) *GetRoomsOutput {
	rooms := uc.tcmClient.GetRooms()

	var domainRooms = make([]domain.Room, 0, len(rooms))
	for _, room := range rooms {
		domainRooms = append(domainRooms, domain.Room{
			ID:          room.ID,
			Name:        room.Name,
			Campus:      domain.Campus(room.Campus),
			PianoType:   domain.PianoType(room.PianoType),
			PianoNumber: room.PianoNumber,
			IsClassroom: room.IsClassroom,
			IsBasement:  room.IsBasement,
			Floor:       room.Floor,
		})
	}

	return &GetRoomsOutput{
		Rooms: domainRooms,
	}
}
