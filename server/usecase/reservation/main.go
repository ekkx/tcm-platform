package reservation

import (
	"context"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/infra/db"
)

type ReservationUsecase interface {
	GetMyReservations(ctx context.Context, input *GetMyReservationsInput) (*GetMyReservationsOutput, error)
}

type ReservationUsecaseImpl struct {
	tcmClient *tcmrsv.Client
	querier   db.Querier
}

func New(tcmClient *tcmrsv.Client, querier db.Querier) ReservationUsecase {
	return &ReservationUsecaseImpl{
		tcmClient: tcmClient,
		querier:   querier,
	}
}

var _ ReservationUsecase = (*ReservationUsecaseImpl)(nil)
