package handler

import (
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/usecase"
	reservation_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/reservation"
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
