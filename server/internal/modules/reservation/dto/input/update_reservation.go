package input

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
	rsv_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
)

type UpdateReservation struct {
	Actor      actor.Actor
	ID         int64
	ExternalID *string
	CampusType enum.CampusType
	Date       time.Time
	FromHour   int32
	FromMinute int32
	ToHour     int32
	ToMinute   int32
	RoomID     string
	BookerName *string
}

func NewUpdateReservation() *UpdateReservation {
	return &UpdateReservation{}
}

func (input *UpdateReservation) FromProto(ctx context.Context, req *rsv_v1.UpdateReservationRequest) *UpdateReservation {
	var date time.Time
	if req.Reservation.Date != nil {
		date = req.Reservation.Date.AsTime()
	}

	input.Actor = ctxhelper.GetActor(ctx)
	input.ID = req.ReservationId
	input.CampusType = enum.CampusType(req.Reservation.CampusType)
	input.Date = date
	input.FromHour = req.Reservation.FromHour
	input.FromMinute = req.Reservation.FromMinute
	input.ToHour = req.Reservation.ToHour
	input.ToMinute = req.Reservation.ToMinute
	input.RoomID = req.Reservation.RoomId
	input.BookerName = req.Reservation.BookerName

	return input
}
