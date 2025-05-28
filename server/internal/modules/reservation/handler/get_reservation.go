package handler

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	rsv_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
)

func (h *Handler) GetReservation(
	ctx context.Context,
	req *rsv_v1.GetReservationRequest,
) (*rsv_v1.GetReservationReply, error) {
	output, err := h.Usecase.GetReservation(ctx, input.NewGetReservation().FromProto(ctx, req))
	if err != nil {
		return nil, err
	}

	return output.ToProto(), nil
}
