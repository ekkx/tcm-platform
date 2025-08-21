package handler

import (
	"github.com/ekkx/tcm-platform/internal/gen/pb/reservation/v1/reservationv1connect"
	"github.com/ekkx/tcm-platform/internal/modules/reservation/usecase"
)

type HandlerImpl struct {
	useCase usecase.UseCase
}

func New(useCase usecase.UseCase) reservationv1connect.ReservationServiceHandler {
	return &HandlerImpl{
		useCase: useCase,
	}
}
