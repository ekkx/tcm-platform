package usecase

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	roomv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/room/v1"
	"github.com/ekkx/tcmrsv-web/server/pkg/ymd"
)

type ListAvailableRoomsInput struct {
	CampusType enum.CampusType
	Date       ymd.YMD
	FromHour   int
	FromMinute int
	ToHour     int
	ToMinute   int
}

func NewListAvailableRoomsInputFromRequest(ctx context.Context, req *connect.Request[roomv1.ListAvailableRoomsRequest]) (*ListAvailableRoomsInput, error) {
	var campusType enum.CampusType
	switch req.Msg.CampusType {
	case roomv1.CampusType_CAMPUS_TYPE_IKEBUKURO:
		campusType = enum.CampusTypeIkebukuro
	case roomv1.CampusType_CAMPUS_TYPE_NAKAMEGURO:
		campusType = enum.CampusTypeNakameguro
	case roomv1.CampusType_CAMPUS_TYPE_UNSPECIFIED:
		campusType = enum.CampusTypeUnknown
	}

	var date ymd.YMD
	date, err := ymd.Parse(req.Msg.Date)
	if err != nil {
		date = ymd.Zero()
	}

	return &ListAvailableRoomsInput{
		CampusType: campusType,
		Date:       date,
		FromHour:   int(req.Msg.FromHour),
		FromMinute: int(req.Msg.FromMinute),
		ToHour:     int(req.Msg.ToHour),
		ToMinute:   int(req.Msg.ToMinute),
	}, nil
}

func (st *ListAvailableRoomsInput) Validate() error {
	if !st.CampusType.IsValid() {
		return errs.ErrInvalidCampusType
	}
	if st.FromHour < 0 || st.FromHour > 23 {
		return errs.ErrInvalidTimeRange
	}
	if st.FromMinute != 0 && st.FromMinute != 30 {
		return errs.ErrInvalidTimeRange
	}
	if st.ToHour < 0 || st.ToHour > 23 {
		return errs.ErrInvalidTimeRange
	}
	if st.ToMinute != 0 && st.ToMinute != 30 {
		return errs.ErrInvalidTimeRange
	}
	if st.FromHour > st.ToHour || (st.FromHour == st.ToHour && st.FromMinute >= st.ToMinute) {
		return errs.ErrInvalidTimeRange
	}
	if st.Date.IsZero() || !st.Date.IsValid() {
		return errs.ErrInvalidDate
	}
	return nil
}
