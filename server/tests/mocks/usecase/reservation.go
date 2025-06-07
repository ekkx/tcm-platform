package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/output"
)

// MockReservationUsecase is a mock implementation of reservation.Usecase
type MockReservationUsecase struct {
	CreateReservationFunc       func(ctx context.Context, input *input.CreateReservation) (*output.CreateReservation, error)
	GetReservationFunc          func(ctx context.Context, input *input.GetReservation) (*output.GetReservation, error)
	GetUserReservationsFunc     func(ctx context.Context, input *input.GetUserReservations) (*output.GetMyReservations, error)
	UpdateReservationFunc       func(ctx context.Context, input *input.UpdateReservation) (*output.UpdateReservation, error)
	DeleteReservationFunc       func(ctx context.Context, input *input.DeleteReservation) error
}

// CreateReservation calls the mock function
func (m *MockReservationUsecase) CreateReservation(ctx context.Context, input *input.CreateReservation) (*output.CreateReservation, error) {
	if m.CreateReservationFunc != nil {
		return m.CreateReservationFunc(ctx, input)
	}
	return nil, nil
}

// GetReservation calls the mock function
func (m *MockReservationUsecase) GetReservation(ctx context.Context, input *input.GetReservation) (*output.GetReservation, error) {
	if m.GetReservationFunc != nil {
		return m.GetReservationFunc(ctx, input)
	}
	return nil, nil
}

// GetUserReservations calls the mock function
func (m *MockReservationUsecase) GetUserReservations(ctx context.Context, input *input.GetUserReservations) (*output.GetMyReservations, error) {
	if m.GetUserReservationsFunc != nil {
		return m.GetUserReservationsFunc(ctx, input)
	}
	return nil, nil
}

// UpdateReservation calls the mock function
func (m *MockReservationUsecase) UpdateReservation(ctx context.Context, input *input.UpdateReservation) (*output.UpdateReservation, error) {
	if m.UpdateReservationFunc != nil {
		return m.UpdateReservationFunc(ctx, input)
	}
	return nil, nil
}

// DeleteReservation calls the mock function
func (m *MockReservationUsecase) DeleteReservation(ctx context.Context, input *input.DeleteReservation) error {
	if m.DeleteReservationFunc != nil {
		return m.DeleteReservationFunc(ctx, input)
	}
	return nil
}