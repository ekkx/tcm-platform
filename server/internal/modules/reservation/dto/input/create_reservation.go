package input

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/core/types"
	reservation_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/reservation"
)

type CreateReservation struct {
	UserID       string
	CampusType   types.CampusType
	Date         *time.Time
	FromHour     int32
	FromMinute   int32
	ToHour       int32
	ToMinute     int32
	IsAutoSelect bool
	RoomID       *string
	BookerName   *string
	PianoNumbers []int32
	PianoTypes   []types.PianoType
	Floors       []int32
	IsBasement   *bool
}

func NewCreateReservation() *CreateReservation {
	return &CreateReservation{}
}

func (input *CreateReservation) Validate() error {
	// Implement validation logic here if needed
	return nil
}

func (input *CreateReservation) FromProto(ctx context.Context, req *reservation_v1.CreateReservationRequest) *CreateReservation {
	var date time.Time
	if req.Reservation.Date != nil {
		date = req.Reservation.Date.AsTime()
	}

	var pianoTypes []types.PianoType
	if req.Reservation.PianoTypes != nil {
		pianoTypes = make([]types.PianoType, len(req.Reservation.PianoTypes))
		for i, pianoType := range req.Reservation.PianoTypes {
			pianoTypes[i] = types.PianoType(pianoType.Number())
		}
	}

	// input.UserID = ctxhelper.GetGRPCRequestUser(ctx)
	input.CampusType = types.CampusType(req.Reservation.Campus.String())
	input.Date = &date
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
