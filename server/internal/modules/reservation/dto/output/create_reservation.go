package output

import (
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	reservation_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
	room_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/room"
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
	case enum.CampusTypeNakameguro:
		campusType = room_v1.CampusType_NAKAMEGURO
	case enum.CampusTypeIkebukuro:
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
