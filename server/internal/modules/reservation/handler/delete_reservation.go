package handler

import (
	"context"

	rsv_v1 "github.com/ekkx/tcmrsv-web/server/internal/api/v1/reservation"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
)

func (h *Handler) DeleteReservation(
	ctx context.Context,
	req *rsv_v1.DeleteReservationRequest,
) (*rsv_v1.DeleteReservationReply, error) {
	err := h.Usecase.DeleteReservation(ctx, input.NewDeleteReservation().FromProto(ctx, req))
	if err != nil {
		return nil, err
	}

	return &rsv_v1.DeleteReservationReply{}, nil
}
