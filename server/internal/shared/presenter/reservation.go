package presenter

import (
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	reservationv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1"
	roomv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/room/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToReservation(rsv *entity.Reservation) *reservationv1.Reservation {
	if rsv == nil {
		return nil
	}

	// TODO: 共通化できる
	var campusType roomv1.CampusType
	switch rsv.CampusType {
	case enum.CampusTypeIkebukuro:
		campusType = roomv1.CampusType_CAMPUS_TYPE_IKEBUKURO
	case enum.CampusTypeNakameguro:
		campusType = roomv1.CampusType_CAMPUS_TYPE_NAKAMEGURO
	default:
		campusType = roomv1.CampusType_CAMPUS_TYPE_UNSPECIFIED
	}

	return &reservationv1.Reservation{
		Id:             rsv.ID.String(),
		OfficialSiteId: rsv.OfficialSiteID,
		User:           ToUser(&rsv.User),
		CampusType:     campusType,
		Room:           ToRoom(&rsv.Room),
		Date:           rsv.Date.String(),
		FromHour:       int32(rsv.FromHour),
		FromMinute:     int32(rsv.FromMinute),
		ToHour:         int32(rsv.ToHour),
		ToMinute:       int32(rsv.ToMinute),
		CreateTime:     timestamppb.New(rsv.CreateTime),
	}
}

func ToReservationList(reservations []*entity.Reservation) []*reservationv1.Reservation {
	if reservations == nil {
		return nil
	}
	result := make([]*reservationv1.Reservation, len(reservations))
	for i, rsv := range reservations {
		result[i] = ToReservation(rsv)
	}
	return result
}
