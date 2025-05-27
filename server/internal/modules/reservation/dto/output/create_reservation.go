package output

import (
	"github.com/ekkx/tcmrsv-web/server/internal/core/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/core/types"
	reservation_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/reservation"
	room_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/room"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreateReservation struct {
	Reservation entity.Reservation
}

func NewCreateReservation(reservation entity.Reservation) *CreateReservation {
	return &CreateReservation{
		Reservation: reservation,
	}
}

func (output *CreateReservation) ToProto() *reservation_v1.CreateReservationReply {
	var campusType room_v1.CampusType
	switch output.Reservation.CampusType {
	case types.CampusTypeNakameguro:
		campusType = room_v1.CampusType_NAKAMEGURO
	case types.CampusTypeIkebukuro:
		campusType = room_v1.CampusType_IKEBUKURO
	default:
		campusType = room_v1.CampusType_CAMPUS_UNSPECIFIED
	}

	return &reservation_v1.CreateReservationReply{
		Reservation: &reservation_v1.Reservation{
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
