package handler

import (
	"context"

	"connectrpc.com/connect"
	reservationv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/reservation/v1"
)

func (h *HandlerImpl) ListReservations(ctx context.Context, req *connect.Request[reservationv1.ListReservationsRequest]) (*connect.Response[reservationv1.ListReservationsResponse], error) {
    return nil, nil
}
