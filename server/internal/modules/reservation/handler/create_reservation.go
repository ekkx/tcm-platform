package handler

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	rsv_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
)

func (h *Handler) CreateReservation(
	ctx context.Context,
	req *rsv_v1.CreateReservationRequest,
) (*rsv_v1.CreateReservationReply, error) {
	output, err := h.Usecase.CreateReservation(ctx, input.NewCreateReservation().FromProto(ctx, req))
	if err != nil {
		return nil, err
	}

	return output.ToProto(), nil
}
