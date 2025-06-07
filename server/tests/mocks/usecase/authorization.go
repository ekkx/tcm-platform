package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/dto/output"
)

// MockAuthorizationUsecase is a mock implementation of authorization.Usecase
type MockAuthorizationUsecase struct {
	AuthorizeFunc   func(ctx context.Context, input *input.Authorize) (*output.Authorize, error)
	ReauthorizeFunc func(ctx context.Context, input *input.Reauthorize) (*output.Reauthorize, error)
}

// Authorize calls the mock function
func (m *MockAuthorizationUsecase) Authorize(ctx context.Context, input *input.Authorize) (*output.Authorize, error) {
	if m.AuthorizeFunc != nil {
		return m.AuthorizeFunc(ctx, input)
	}
	return nil, nil
}

// Reauthorize calls the mock function
func (m *MockAuthorizationUsecase) Reauthorize(ctx context.Context, input *input.Reauthorize) (*output.Reauthorize, error) {
	if m.ReauthorizeFunc != nil {
		return m.ReauthorizeFunc(ctx, input)
	}
	return nil, nil
}