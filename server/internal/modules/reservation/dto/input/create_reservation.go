package input

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	reservation_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
)

type CreateReservation struct {
	UserID     string          `validate:"required"`
	CampusType enum.CampusType `validate:"required"`
	Date       time.Time       `validate:"required"`
	FromHour   int             `validate:"gte=0,lte=23"`
	FromMinute int             `validate:"oneof=0 30"`
	ToHour     int             `validate:"gte=0,lte=23"`
	ToMinute   int             `validate:"oneof=0 30"`
	RoomID     string          `validate:"required"`
	BookerName *string         `validate:"omitempty"`
}

func NewCreateReservation() *CreateReservation {
	return &CreateReservation{}
}

func (input *CreateReservation) Validate() error {
	if !input.CampusType.IsValid() {
		return errs.ErrInvalidCampusType
	}
	if input.FromHour > input.ToHour || (input.FromHour == input.ToHour && input.FromMinute >= input.ToMinute) {
		return errs.ErrInvalidTimeRange
	}
	return validate.Struct(input)
}

func (input *CreateReservation) FromProto(ctx context.Context, req *reservation_v1.CreateReservationRequest) *CreateReservation {
	var date time.Time
	if req.Reservation.Date != nil {
		date = req.Reservation.Date.AsTime()
	}

	input.UserID = ctxhelper.GetActor(ctx).ID
	input.CampusType = enum.CampusType(req.Reservation.CampusType)
	input.Date = date
	input.FromHour = int(req.Reservation.FromHour)
	input.FromMinute = int(req.Reservation.FromMinute)
	input.ToHour = int(req.Reservation.ToHour)
	input.ToMinute = int(req.Reservation.ToMinute)
	input.RoomID = req.Reservation.RoomId
	input.BookerName = req.Reservation.BookerName

	return input
}
