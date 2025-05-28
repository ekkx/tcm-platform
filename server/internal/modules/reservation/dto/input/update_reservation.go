package input

import (
	"context"
	"time"

	rsv_v1 "github.com/ekkx/tcmrsv-web/server/internal/api/v1/reservation"
	"github.com/ekkx/tcmrsv-web/server/internal/core/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
)

type UpdateReservation struct {
	Actor        actor.Actor
	ID           int64
	ExternalID   *string
	CampusType   enum.CampusType
	Date         time.Time
	FromHour     int32
	FromMinute   int32
	ToHour       int32
	ToMinute     int32
	IsAutoSelect bool
	RoomID       *string
	BookerName   *string
	PianoNumbers []int32
	PianoTypes   []enum.PianoType
	Floors       []int32
	IsBasement   *bool
}

func NewUpdateReservation() *UpdateReservation {
	return &UpdateReservation{}
}

func (input *UpdateReservation) FromProto(ctx context.Context, req *rsv_v1.UpdateReservationRequest) *UpdateReservation {
	var date time.Time
	if req.Reservation.Date != nil {
		date = req.Reservation.Date.AsTime()
	}

	var pianoTypes []enum.PianoType
	if req.Reservation.PianoTypes != nil {
		pianoTypes = make([]enum.PianoType, len(req.Reservation.PianoTypes))
		for i, pianoType := range req.Reservation.PianoTypes {
			pianoTypes[i] = enum.PianoType(pianoType)
		}
	}

	input.Actor = ctxhelper.GetActor(ctx)
	input.ID = req.ReservationId
	input.CampusType = enum.CampusType(req.Reservation.CampusType)
	input.Date = date
	input.FromHour = req.Reservation.FromHour
	input.FromMinute = req.Reservation.FromMinute
	input.ToHour = req.Reservation.ToHour
	input.ToMinute = req.Reservation.ToMinute
	input.IsAutoSelect = req.Reservation.IsAutoSelect
	input.RoomID = req.Reservation.RoomId
	input.BookerName = req.Reservation.BookerName
	input.PianoNumbers = req.Reservation.PianoNumbers
	input.PianoTypes = pianoTypes
	input.Floors = req.Reservation.Floors
	input.IsBasement = req.Reservation.IsBasement

	return input
}
