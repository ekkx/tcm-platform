package usecase

import (
	"context"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/output"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/utils"
)

func (u *Usecase) GetUserReservations(ctx context.Context, params *input.GetUserReservations) (*output.GetMyReservations, error) {
	if err := params.Validate(); err != nil {
		return nil, errs.ErrInvalidArgument.WithCause(err)
	}

	// 日付未指定の場合は今日の日付を基準にする
	var fromDate time.Time
	if params.FromDate.IsZero() {
		now := time.Now()
		fromDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, utils.JST())
	} else {
		fromDate = params.FromDate.In(utils.JST())
	}

	rsvs, err := u.rsvRepo.GetUserReservations(ctx, &repository.GetUserReservationsArgs{
		UserID:   params.UserID,
		FromDate: fromDate,
	})
	if err != nil {
		return nil, errs.ErrInternal.WithCause(err)
	}

	return output.NewGetMyReservations(rsvs), nil
}
