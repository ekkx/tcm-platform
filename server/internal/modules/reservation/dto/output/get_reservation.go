package output

import (
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	rsv_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
	room_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/room"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetReservation struct {
	Reservation entity.Reservation
}

func NewGetReservation(reservation entity.Reservation) *GetReservation {
	return &GetReservation{
		Reservation: reservation,
	}
}

func (output *GetReservation) ToProto() *rsv_v1.GetReservationReply {
	var campusType room_v1.CampusType
	switch output.Reservation.CampusType {
	case enum.CampusTypeNakameguro:
		campusType = room_v1.CampusType_NAKAMEGURO
	case enum.CampusTypeIkebukuro:
		campusType = room_v1.CampusType_IKEBUKURO
	default:
		campusType = room_v1.CampusType_CAMPUS_UNSPECIFIED
	}

	// For date-only fields, we want to preserve the date as-is
	// Create a UTC time with the same date components
	dateUTC := time.Date(output.Reservation.Date.Year(), output.Reservation.Date.Month(), output.Reservation.Date.Day(), 0, 0, 0, 0, time.UTC)
	
	return &rsv_v1.GetReservationReply{
		Reservation: &rsv_v1.Reservation{
			Id:         int64(output.Reservation.ID),
			ExternalId: output.Reservation.ExternalID,
			CampusType: campusType,
			Date:       timestamppb.New(dateUTC),
			RoomId:     output.Reservation.RoomID,
			FromHour:   int32(output.Reservation.FromHour),
			FromMinute: int32(output.Reservation.FromMinute),
			ToHour:     int32(output.Reservation.ToHour),
			ToMinute:   int32(output.Reservation.ToMinute),
			BookerName: output.Reservation.BookerName,
			CreatedAt:  timestamppb.New(output.Reservation.CreatedAt.In(time.UTC)),
		},
	}
}
