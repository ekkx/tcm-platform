package usecase

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	reservationv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1"
	"github.com/ekkx/tcmrsv-web/server/pkg/actor"
	"github.com/ekkx/tcmrsv-web/server/pkg/ymd"
)

type CreateReservationInput struct {
	Actor      actor.Actor
	CampusType enum.CampusType
	Date       ymd.YMD
	FromHour   int
	FromMinute int
	ToHour     int
	ToMinute   int
	RoomID     string
}

func NewCreateReservationInputFromRequest(ctx context.Context, req *connect.Request[reservationv1.CreateReservationRequest]) (*CreateReservationInput, error) {
	st := &CreateReservationInput{}

	actor := ctxhelper.Actor(ctx)
	if actor == nil {
		return nil, errs.ErrUnauthorized
	}
	st.Actor = *actor

	// TODO: ここら辺共通化できる
	var campusType enum.CampusType
	switch req.Msg.CampusType {
	case reservationv1.CampusType_CAMPUS_TYPE_IKEBUKURO:
		campusType = enum.CampusTypeIkebukuro
	case reservationv1.CampusType_CAMPUS_TYPE_NAKAMEGURO:
		campusType = enum.CampusTypeNakameguro
	case reservationv1.CampusType_CAMPUS_TYPE_UNSPECIFIED:
		campusType = enum.CampusTypeUnknown
	}

	var date ymd.YMD
	date, err := ymd.Parse(req.Msg.Date)
	if err != nil {
		date = ymd.Zero()
	}

	st.CampusType = campusType
	st.Date = date
	st.FromHour = int(req.Msg.FromHour)
	st.FromMinute = int(req.Msg.FromMinute)
	st.ToHour = int(req.Msg.ToHour)
	st.ToMinute = int(req.Msg.ToMinute)
	st.RoomID = req.Msg.RoomId

	return st, nil
}

func (st *CreateReservationInput) Validate() error {
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
	if st.Date.Before(ymd.Today().AddDays(3)) {
		return errs.ErrReservationTooSoon
	}
	return nil
}
