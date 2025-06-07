package input

import (
	"context"
	"time"

	rsv_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/utils"
)

type GetUserReservations struct {
	UserID   string
	FromDate time.Time
}

func NewGetUserReservations() *GetUserReservations {
	return &GetUserReservations{}
}

func (input *GetUserReservations) Validate() error {
	// 日付が今日より前の場合はエラー
	if !input.FromDate.IsZero() && input.FromDate.Before(time.Now().In(utils.JST())) {
		return errs.ErrDateMustBeTodayOrFuture
	}
	return nil
}

func (input *GetUserReservations) FromProto(ctx context.Context, req *rsv_v1.GetUserReservationsRequest) *GetUserReservations {
	input.UserID = ctxhelper.GetActor(ctx).ID
	input.FromDate = req.FromDate.AsTime()
	return input
}
