package handler

import (
	reservation_v1 "github.com/ekkx/tcmrsv-web/server/internal/api/v1/reservation"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/usecase"
)

type Handler struct {
	reservation_v1.UnimplementedReservationServiceServer

	Usecase *usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}
