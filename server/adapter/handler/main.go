package handler

import (
	"github.com/ekkx/tcmrsv-web/server/adapter/api"
	"github.com/ekkx/tcmrsv-web/server/usecase/authorize"
	"github.com/ekkx/tcmrsv-web/server/usecase/reservation"
	"github.com/ekkx/tcmrsv-web/server/usecase/room"
)

type Handler struct {
	authorizeUsecase   authorize.AuthorizeUsecase
	reservationUsecase reservation.ReservationUsecase
	roomUsecase        room.RoomUsecase
}

func New(
	authorizeUsecase authorize.AuthorizeUsecase,
	reservationUsecase reservation.ReservationUsecase,
	roomUsecase room.RoomUsecase,
) *Handler {
	return &Handler{
		authorizeUsecase:   authorizeUsecase,
		reservationUsecase: reservationUsecase,
		roomUsecase:        roomUsecase,
	}
}

var _ api.ServerInterface = (*Handler)(nil)
