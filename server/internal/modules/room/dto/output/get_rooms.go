package output

import (
	room_v1 "github.com/ekkx/tcmrsv-web/server/internal/api/v1/room"
	"github.com/ekkx/tcmrsv-web/server/internal/core/entity"
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
	reply := &room_v1.GetRoomsReply{
		Rooms: make([]*room_v1.Room, len(output.Rooms)),
	}

	for i, room := range output.Rooms {
		reply.Rooms[i] = &room_v1.Room{
			Id:          room.ID,
			Name:        room.Name,
			PianoType:   room_v1.PianoType(room.PianoType),
			PianoNumber: room.PianoNumber,
			IsClassroom: room.IsClassroom,
			IsBasement:  room.IsBasement,
			CampusType:  room_v1.CampusType(room.CampusType),
			Floor:       room.Floor,
		}
	}

	return reply
}
