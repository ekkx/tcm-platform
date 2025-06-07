package input

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
	rsv_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
)

type UpdateReservation struct {
	Actor      actor.Actor
	ID         int             `validate:"required"`
	ExternalID *string         `validate:"omitempty"`
	CampusType enum.CampusType `validate:"required"`
	Date       time.Time       `validate:"required"`
	FromHour   int             `validate:"gte=0,lte=23"`
	FromMinute int             `validate:"oneof=0 30"`
	ToHour     int             `validate:"gte=0,lte=23"`
	ToMinute   int             `validate:"oneof=0 30"`
	RoomID     string          `validate:"required"`
	BookerName *string         `validate:"omitempty"`
}

func NewUpdateReservation() *UpdateReservation {
	return &UpdateReservation{}
}

func (input *UpdateReservation) Validate() error {
	if !input.CampusType.IsValid() {
		return errs.ErrInvalidCampusType
	}
	if input.FromHour > input.ToHour || (input.FromHour == input.ToHour && input.FromMinute >= input.ToMinute) {
		return errs.ErrInvalidTimeRange
	}
	return validate.Struct(input)
}

func (input *UpdateReservation) FromProto(ctx context.Context, req *rsv_v1.UpdateReservationRequest) *UpdateReservation {
	var date time.Time
	if req.Reservation.Date != nil {
		date = req.Reservation.Date.AsTime()
	}

	input.Actor = ctxhelper.GetActor(ctx)
	input.ID = int(req.ReservationId)
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
