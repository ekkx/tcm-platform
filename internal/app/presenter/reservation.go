package presenter

import (
	"github.com/ekkx/tcm-platform/internal/app/assemble"
	"github.com/ekkx/tcm-platform/internal/domain/valueobject"
	reservationv1 "github.com/ekkx/tcm-platform/internal/gen/pb/reservation/v1"
	roomv1 "github.com/ekkx/tcm-platform/internal/gen/pb/room/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToReservation(v *assemble.ReservationView) *reservationv1.Reservation {
	if v == nil {
		return nil
	}

	// TODO: 共通化できる
	var campusType roomv1.CampusType
	switch v.Reservation.CampusType {
	case valueobject.CampusTypeIkebukuro:
		campusType = roomv1.CampusType_CAMPUS_TYPE_IKEBUKURO
	case valueobject.CampusTypeNakameguro:
		campusType = roomv1.CampusType_CAMPUS_TYPE_NAKAMEGURO
	default:
		campusType = roomv1.CampusType_CAMPUS_TYPE_UNSPECIFIED
	}

	return &reservationv1.Reservation{
		Id:             v.Reservation.ID.String(),
		OfficialSiteId: v.Reservation.OfficialSiteID,
		User:           ToUser(&v.UserView),
		CampusType:     campusType,
		Room:           ToRoom(&v.Room),
		Date:           v.Reservation.Date.String(),
		FromHour:       int32(v.Reservation.FromHour),
		FromMinute:     int32(v.Reservation.FromMinute),
		ToHour:         int32(v.Reservation.ToHour),
		ToMinute:       int32(v.Reservation.ToMinute),
		CreateTime:     timestamppb.New(v.Reservation.CreateTime),
	}
}

func ToReservationList(v []*assemble.ReservationView) []*reservationv1.Reservation {
	if v == nil {
		return nil
	}
	result := make([]*reservationv1.Reservation, len(v))
	for i, rsv := range v {
		result[i] = ToReservation(rsv)
	}
	return result
}
