package handler_test

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/output"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
)

// ptr is a helper function to get a pointer to a string
func ptr(s string) *string {
	return &s
}

// testReservationHandler is a test handler implementation
type testReservationHandler struct {
	reservation.UnimplementedReservationServiceServer
	mockCreateReservation   func(context.Context, *input.CreateReservation) (*output.CreateReservation, error)
	mockGetReservation      func(context.Context, *input.GetReservation) (*output.GetReservation, error)
	mockGetUserReservations func(context.Context, *input.GetUserReservations) (*output.GetMyReservations, error)
	mockUpdateReservation   func(context.Context, *input.UpdateReservation) (*output.UpdateReservation, error)
	mockDeleteReservation   func(context.Context, *input.DeleteReservation) error
}

func (h *testReservationHandler) CreateReservation(ctx context.Context, req *reservation.CreateReservationRequest) (*reservation.CreateReservationReply, error) {
	if h.mockCreateReservation != nil {
		output, err := h.mockCreateReservation(ctx, input.NewCreateReservation().FromProto(ctx, req))
		if err != nil {
			return nil, err
		}
		return output.ToProto(), nil
	}
	return &reservation.CreateReservationReply{}, nil
}

func (h *testReservationHandler) GetReservation(ctx context.Context, req *reservation.GetReservationRequest) (*reservation.GetReservationReply, error) {
	if h.mockGetReservation != nil {
		output, err := h.mockGetReservation(ctx, input.NewGetReservation().FromProto(ctx, req))
		if err != nil {
			return nil, err
		}
		return output.ToProto(), nil
	}
	return &reservation.GetReservationReply{}, nil
}

func (h *testReservationHandler) GetUserReservations(ctx context.Context, req *reservation.GetUserReservationsRequest) (*reservation.GetUserReservationsReply, error) {
	if h.mockGetUserReservations != nil {
		output, err := h.mockGetUserReservations(ctx, input.NewGetUserReservations().FromProto(ctx, req))
		if err != nil {
			return nil, err
		}
		return output.ToProto(), nil
	}
	return &reservation.GetUserReservationsReply{}, nil
}

func (h *testReservationHandler) GetMyReservations(ctx context.Context, req *reservation.GetUserReservationsRequest) (*reservation.GetUserReservationsReply, error) {
	return h.GetUserReservations(ctx, req)
}

func (h *testReservationHandler) UpdateReservation(ctx context.Context, req *reservation.UpdateReservationRequest) (*reservation.UpdateReservationReply, error) {
	if h.mockUpdateReservation != nil {
		output, err := h.mockUpdateReservation(ctx, input.NewUpdateReservation().FromProto(ctx, req))
		if err != nil {
			return nil, err
		}
		return output.ToProto(), nil
	}
	return &reservation.UpdateReservationReply{}, nil
}

func (h *testReservationHandler) DeleteReservation(ctx context.Context, req *reservation.DeleteReservationRequest) (*reservation.DeleteReservationReply, error) {
	if h.mockDeleteReservation != nil {
		err := h.mockDeleteReservation(ctx, input.NewDeleteReservation().FromProto(ctx, req))
		if err != nil {
			return nil, err
		}
	}
	return &reservation.DeleteReservationReply{}, nil
}
