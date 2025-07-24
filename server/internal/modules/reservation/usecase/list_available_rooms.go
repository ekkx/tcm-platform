package usecase

import (
	"context"
	"slices"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/mapper"
)

func (uc *UseCaseImpl) ListAvailableRooms(ctx context.Context, input *ListAvailableRoomsInput) (*ListAvailableRoomsOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	roomIDs, err := uc.reservationRepo.ListUnavailableRoomIDs(ctx, &repository.ListUnavailableRoomIDsParams{
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
	case enum.CampusTypeIkebukuro:
		tcmCampus = tcmrsv.CampusIkebukuro
	case enum.CampusTypeNakameguro:
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
