package handler

import (
	"context"

	"connectrpc.com/connect"
	reservationv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1"
)

func (h *Handler) GetReservation(ctx context.Context, req *connect.Request[reservationv1.GetReservationRequest]) (*connect.Response[reservationv1.GetReservationResponse], error) {
    return nil, nil
}
