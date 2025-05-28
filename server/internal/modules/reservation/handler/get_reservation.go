package handler

import (
	"context"

	rsv_v1 "github.com/ekkx/tcmrsv-web/server/internal/api/v1/reservation"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
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
