package output

import (
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	rsv_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
	room_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/room"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreateReservation struct {
	Reservations []entity.Reservation
}

func NewCreateReservation(reservations []entity.Reservation) *CreateReservation {
	return &CreateReservation{
		Reservations: reservations,
	}
}

func (output *CreateReservation) ToProto() *rsv_v1.CreateReservationReply {
	protoRsvs := make([]*rsv_v1.Reservation, len(output.Reservations))
	if len(output.Reservations) == 0 {
		return &rsv_v1.CreateReservationReply{
			Reservations: protoRsvs,
		}
	}

	for _, rsv := range output.Reservations {
		var campusType room_v1.CampusType
		switch rsv.CampusType {
		case enum.CampusTypeNakameguro:
			campusType = room_v1.CampusType_NAKAMEGURO
		case enum.CampusTypeIkebukuro:
			campusType = room_v1.CampusType_IKEBUKURO
		default:
			campusType = room_v1.CampusType_CAMPUS_UNSPECIFIED
		}

		protoRsvs = append(protoRsvs, &rsv_v1.Reservation{
			Id:         int64(rsv.ID),
			ExternalId: rsv.ExternalID,
			CampusType: campusType,
			Date:       timestamppb.New(rsv.Date),
			RoomId:     rsv.RoomID,
			FromHour:   int32(rsv.FromHour),
			FromMinute: int32(rsv.FromMinute),
			ToHour:     int32(rsv.ToHour),
			ToMinute:   int32(rsv.ToMinute),
			BookerName: rsv.BookerName,
			CreatedAt:  timestamppb.New(rsv.CreatedAt),
		})
	}

	return &rsv_v1.CreateReservationReply{
		Reservations: protoRsvs,
	}
}
