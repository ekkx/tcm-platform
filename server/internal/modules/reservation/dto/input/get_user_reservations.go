package input

import (
	"context"
	"time"

	rsv_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
)

type GetUserReservations struct {
	UserID string
	Date   time.Time
}

func NewGetUserReservations() *GetUserReservations {
	return &GetUserReservations{}
}

func (input *GetUserReservations) FromProto(ctx context.Context, req *rsv_v1.GetUserReservationsRequest) *GetUserReservations {
	input.UserID = ctxhelper.GetActor(ctx).ID
	// TODO: 日付も追加する？
	return input
}
