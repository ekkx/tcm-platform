package output

import (
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	room_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/room"
)

type GetRooms struct {
	Rooms []entity.Room
}

func NewGetRooms(rooms []entity.Room) *GetRooms {
	return &GetRooms{
		Rooms: rooms,
	}
}

func (output *GetRooms) ToProto() *room_v1.GetRoomsReply {
	protoRooms := make([]*room_v1.Room, 0, len(output.Rooms))

	for _, room := range output.Rooms {
		protoRooms = append(protoRooms, &room_v1.Room{
			Id:          room.ID,
			Name:        room.Name,
			PianoType:   room_v1.PianoType(room.PianoType),
			PianoNumber: int32(room.PianoNumber),
			IsClassroom: room.IsClassroom,
			IsBasement:  room.IsBasement,
			CampusType:  room_v1.CampusType(room.CampusType),
			Floor:       int32(room.Floor),
		})
	}

	return &room_v1.GetRoomsReply{
		Rooms: protoRooms,
	}
}
