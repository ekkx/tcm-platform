package handler

import (
	"github.com/ekkx/tcmrsv-web/server/adapter/api"
	"github.com/ekkx/tcmrsv-web/server/usecase/authorize"
	"github.com/ekkx/tcmrsv-web/server/usecase/reservation"
)

type Handler struct {
	authorizeUsecase   authorize.AuthorizeUsecase
	reservationUsecase reservation.ReservationUsecase
}

func New(
	authorizeUsecase authorize.AuthorizeUsecase,
	reservationUsecase reservation.ReservationUsecase,
) *Handler {
	return &Handler{
		authorizeUsecase:   authorizeUsecase,
		reservationUsecase: reservationUsecase,
	}
}

var _ api.ServerInterface = (*Handler)(nil)
