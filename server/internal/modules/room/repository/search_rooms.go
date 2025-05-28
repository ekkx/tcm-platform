package repository

import (
	"context"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/tcmutil"
)

type SearchRoomsArgs struct {
	Name         *string
	ID           *string
	PianoNumbers []int32
	PianoTypes   []enum.PianoType
	Floors       []int32
	IsBasement   *bool
	Campuses     []enum.CampusType
}

func (repo *Repository) SearchRooms(ctx context.Context, args SearchRoomsArgs) []entity.Room {
	var pianoNumbers []int
	for _, n := range args.PianoNumbers {
		pianoNumbers = append(pianoNumbers, int(n))
	}

	var pianoTypes []tcmrsv.RoomPianoType
	for _, pt := range args.PianoTypes {
		pianoTypes = append(pianoTypes, tcmrsv.RoomPianoType(pt))
	}

	var floors []int
	for _, f := range args.Floors {
		floors = append(floors, int(f))
	}

	var campuses []tcmrsv.Campus
	for _, c := range args.Campuses {
		campuses = append(campuses, tcmrsv.Campus(c))
	}

	tcmRooms := repo.tcmClient.GetRoomsFiltered(tcmrsv.GetRoomsFilteredParams{
		Name:         args.Name,
		ID:           args.ID,
		PianoNumbers: pianoNumbers,
		PianoTypes:   pianoTypes,
		Floors:       floors,
		IsBasement:   args.IsBasement,
		Campuses:     campuses,
	})

	rooms := make([]entity.Room, len(tcmRooms))
	for i, tcmRoom := range tcmRooms {
		rooms[i] = entity.Room{
			ID:          tcmRoom.ID,
			Name:        tcmRoom.Name,
			PianoType:   tcmutil.ConvertRoomPianoType(tcmRoom.PianoType),
			PianoNumber: int32(tcmRoom.PianoNumber),
			IsClassroom: tcmRoom.IsClassroom,
			IsBasement:  tcmRoom.IsBasement,
			CampusType:  tcmutil.ConvertCampusType(tcmRoom.Campus),
			Floor:       int32(tcmRoom.Floor),
		}
	}

	return rooms
}
