package usecase

import (
	"context"
	"slices"

	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/domain/valueobject"
	"github.com/ekkx/tcm-platform/internal/modules/room/mapper"
	"github.com/ekkx/tcm-platform/internal/modules/room/repository"
	"github.com/ekkx/tcm-platform/internal/platform/errs"
	"github.com/ekkx/tcmrsv"
)

func (uc *UseCaseImpl) ListAvailableRooms(ctx context.Context, input *ListAvailableRoomsInput) (*ListAvailableRoomsOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	roomIDs, err := uc.roomRepo.ListUnavailableRoomIDs(ctx, &repository.ListUnavailableRoomIDsParams{
		CampusType: input.CampusType,
		Date:       input.Date,
		FromHour:   input.FromHour,
		FromMinute: input.FromMinute,
		ToHour:     input.ToHour,
		ToMinute:   input.ToMinute,
	})
	if err != nil {
		return nil, err
	}

	var tcmCampus tcmrsv.Campus
	switch input.CampusType {
	case valueobject.CampusTypeIkebukuro:
		tcmCampus = tcmrsv.CampusIkebukuro
	case valueobject.CampusTypeNakameguro:
		tcmCampus = tcmrsv.CampusNakameguro
	default:
		return nil, errs.ErrInvalidCampusType
	}

	rooms := tcmrsv.New().GetRoomsFiltered(tcmrsv.GetRoomsFilteredParams{
		Campuses: []tcmrsv.Campus{tcmCampus},
	})

	var availableRooms []*entity.Room
	for _, room := range rooms {
		if !slices.Contains(roomIDs, room.ID) {
			availableRooms = append(availableRooms, mapper.ToRoom(&room))
		}
	}

	return NewListAvailableRoomsOutput(availableRooms), nil
}
