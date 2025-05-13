package presenter

import (
	"github.com/ekkx/tcmrsv-web/server/adapter/api"
	"github.com/ekkx/tcmrsv-web/server/domain"
)

func toRoom(v *domain.Room) *api.Room {
	if v == nil {
		return nil
	}

	return &api.Room{
		CampusCode:  api.RoomCampusCode(v.Campus),
		Floor:       float32(v.Floor),
		Id:          v.ID,
		IsBasement:  v.IsBasement,
		IsClassroom: v.IsClassroom,
		Name:        v.Name,
		PianoNumber: v.PianoNumber,
		PianoType:   api.RoomPianoType(v.PianoType),
	}
}

func toRoomList(v *[]domain.Room) []api.Room {
	if v == nil {
		return []api.Room{}
	}

	list := make([]api.Room, 0, len(*v))
	for _, room := range *v {
		if converted := toRoom(&room); converted != nil {
			list = append(list, *converted)
		}
	}
	return list
}
