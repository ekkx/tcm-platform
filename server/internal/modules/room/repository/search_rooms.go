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
	PianoNumbers []int
	PianoTypes   []enum.PianoType
	Floors       []int
	IsBasement   *bool
	CampusTypes  []enum.CampusType
}

func (repo *Repository) SearchRooms(ctx context.Context, args *SearchRoomsArgs) []entity.Room {
	var pianoTypes []tcmrsv.RoomPianoType
	for _, pt := range args.PianoTypes {
		pianoTypes = append(pianoTypes, tcmutil.ToTCMRoomPianoType(pt))
	}

	var campuses []tcmrsv.Campus
	for _, c := range args.CampusTypes {
		campuses = append(campuses, tcmutil.ToTCMCampusType(c))
	}

	tcmRooms := repo.tcmClient.GetRoomsFiltered(tcmrsv.GetRoomsFilteredParams{
		Name:         args.Name,
		ID:           args.ID,
		PianoNumbers: args.PianoNumbers,
		PianoTypes:   pianoTypes,
		Floors:       args.Floors,
		IsBasement:   args.IsBasement,
		Campuses:     campuses,
	})

	rooms := make([]entity.Room, len(tcmRooms))
	for i, tcmRoom := range tcmRooms {
		rooms[i] = entity.Room{
			ID:          tcmRoom.ID,
			Name:        tcmRoom.Name,
			PianoType:   tcmutil.ToDomainPianoType(tcmRoom.PianoType),
			PianoNumber: tcmRoom.PianoNumber,
			IsClassroom: tcmRoom.IsClassroom,
			IsBasement:  tcmRoom.IsBasement,
			CampusType:  tcmutil.ToDomainCampusType(tcmRoom.Campus),
			Floor:       tcmRoom.Floor,
		}
	}

	return rooms
}
