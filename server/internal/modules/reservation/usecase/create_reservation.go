package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/output"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
)

func (u *Usecase) CreateReservation(ctx context.Context, params *input.CreateReservation) (*output.CreateReservation, error) {
	// TODO: if params.IsAutoSelect

	rsv, err := u.rsvrepo.CreateReservation(ctx, &repository.CreateReservationArgs{
		UserID:     params.UserID,
		CampusType: params.CampusType,
		RoomID:     *params.RoomID,
		Date:       *params.Date,
		FromHour:   params.FromHour,
		FromMinute: params.FromMinute,
		ToHour:     params.ToHour,
		ToMinute:   params.ToMinute,
		BookerName: params.BookerName,
	})
	if err != nil {
		return nil, err
	}

	return output.NewCreateReservation(rsv), nil
}
