package output

import (
	rsv_v1 "github.com/ekkx/tcmrsv-web/server/internal/api/v1/reservation"
	room_v1 "github.com/ekkx/tcmrsv-web/server/internal/api/v1/room"
	"github.com/ekkx/tcmrsv-web/server/internal/core/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/core/enum"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UpdateReservation struct {
	Reservation entity.Reservation
}

func NewUpdateReservation(reservation entity.Reservation) *UpdateReservation {
	return &UpdateReservation{
		Reservation: reservation,
	}
}

func (output *UpdateReservation) ToProto() *rsv_v1.UpdateReservationReply {
	var campusType room_v1.CampusType
	switch output.Reservation.CampusType {
	case enum.CampusTypeNakameguro:
		campusType = room_v1.CampusType_NAKAMEGURO
	case enum.CampusTypeIkebukuro:
		campusType = room_v1.CampusType_IKEBUKURO
	default:
		campusType = room_v1.CampusType_CAMPUS_UNSPECIFIED
	}

	return &rsv_v1.UpdateReservationReply{
		Reservation: &rsv_v1.Reservation{
			Id:         output.Reservation.ID,
			ExternalId: output.Reservation.ExternalID,
			CampusType: campusType,
			Date:       timestamppb.New(output.Reservation.Date),
			RoomId:     output.Reservation.RoomID,
			FromHour:   output.Reservation.FromHour,
			FromMinute: output.Reservation.FromMinute,
			ToHour:     output.Reservation.ToHour,
			ToMinute:   output.Reservation.ToMinute,
			BookerName: output.Reservation.BookerName,
			CreatedAt:  timestamppb.New(output.Reservation.CreatedAt),
		},
	}
}
