package handler

import (
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/usecase"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1/reservationv1connect"
)

type HandlerImpl struct {
	useCase usecase.UseCase
}

func New(useCase usecase.UseCase) reservationv1connect.ReservationServiceHandler {
	return &HandlerImpl{
		useCase: useCase,
	}
}
