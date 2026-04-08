package mapper

import (
	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/domain/valueobject"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
)

func ToReservation(rsv *sqlc.Reservation) *entity.Reservation {
	var campusType valueobject.CampusType
	switch rsv.CampusType {
	case sqlc.CampusTypeIkebukuro:
		campusType = valueobject.CampusTypeIkebukuro
	case sqlc.CampusTypeNakameguro:
		campusType = valueobject.CampusTypeNakameguro
	default:
		campusType = valueobject.CampusTypeUnknown
	}

	var status valueobject.ReservationStatus
	switch rsv.Status {
	case sqlc.ReservationStatusSuccess:
		status = valueobject.ReservationStatusSuccess
	case sqlc.ReservationStatusFailed:
		status = valueobject.ReservationStatusFailed
	default:
		status = valueobject.ReservationStatusPending
	}

	return &entity.Reservation{
		ID:             rsv.ID,
		OfficialSiteID: rsv.OfficialSiteID,
		UserID:         rsv.UserID,
		CampusType:     campusType,
		RoomID:         rsv.RoomID,
		Date:           rsv.Date,
		FromHour:       int(rsv.FromHour),
		FromMinute:     int(rsv.FromMinute),
		ToHour:         int(rsv.ToHour),
		ToMinute:       int(rsv.ToMinute),
		Status:         status,
		Note:           rsv.Note,
		CreateTime:     rsv.CreateTime,
	}
}
