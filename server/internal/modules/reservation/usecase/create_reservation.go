package usecase

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/core/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/core/vo"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
)

type CreateReservationInput struct {
	UserID       string
	Campus       vo.Campus
	Date         *time.Time
	FromHour     int32
	FromMinute   int32
	ToHour       int32
	ToMinute     int32
	IsAutoSelect bool
	RoomId       *string
	BookerName   *string
	PianoNumbers []int32
	PianoTypes   []vo.PianoType
	Floors       []int32
	IsBasement   *bool
}

type CreateReservationOutput struct {
	Reservation entity.Reservation
}

func (u *Usecase) CreateReservation(ctx context.Context, input *CreateReservationInput) (*CreateReservationOutput, error) {
	var rsv entity.Reservation

	if input.IsAutoSelect {
		return &CreateReservationOutput{
			Reservation: rsv,
		}, nil
	} else {
		rsv, err := u.rsvrepo.CreateReservation(ctx, &repository.CreateReservationArgs{
			UserID:     input.UserID,
			Campus:     input.Campus,
			RoomID:     *input.RoomId,
			Date:       *input.Date,
			FromHour:   input.FromHour,
			FromMinute: input.FromMinute,
			ToHour:     input.ToHour,
			ToMinute:   input.ToMinute,
			BookerName: input.BookerName,
		})
		if err != nil {
			return nil, err
		}

		return &CreateReservationOutput{
			Reservation: rsv,
		}, nil
	}
}
