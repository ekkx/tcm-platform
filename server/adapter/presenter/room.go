package presenter

import (
	"github.com/ekkx/tcmrsv-web/server/adapter/api"
	"github.com/ekkx/tcmrsv-web/server/usecase/room"
)

func GetRooms(output *room.GetRoomsOutput) *api.RoomList {
	return &api.RoomList{
		Rooms: toRoomList(&output.Rooms),
	}
}
