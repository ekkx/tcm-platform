package handler

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	reservation_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/reservation"
)

func (h *Handler) CreateReservation(
	ctx context.Context,
	req *reservation_v1.CreateReservationRequest,
) (*reservation_v1.CreateReservationReply, error) {
	output, err := h.Usecase.CreateReservation(ctx, input.NewCreateReservation().FromProto(ctx, req))
	if err != nil {
		return nil, err
	}

	return output.ToProto(), nil
}
