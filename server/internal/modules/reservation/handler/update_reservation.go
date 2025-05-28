package handler

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	rsv_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
)

func (h *Handler) UpdateReservation(
	ctx context.Context,
	req *rsv_v1.UpdateReservationRequest,
) (*rsv_v1.UpdateReservationReply, error) {
	output, err := h.Usecase.UpdateReservation(ctx, input.NewUpdateReservation().FromProto(ctx, req))
	if err != nil {
		return nil, err
	}

	return output.ToProto(), nil
}
