package handler

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	rsv_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
)

func (h *Handler) GetUserReservations(
	ctx context.Context,
	req *rsv_v1.GetUserReservationsRequest,
) (*rsv_v1.GetUserReservationsReply, error) {
	output, err := h.Usecase.GetUserReservations(ctx, input.NewGetUserReservations().FromProto(ctx, req))
	if err != nil {
		return nil, err
	}

	return output.ToProto(), nil
}
