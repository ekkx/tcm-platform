package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/room/dto/output"
)

// MockRoomUsecase is a mock implementation of room.Usecase
type MockRoomUsecase struct {
	GetRoomsFunc func(ctx context.Context) *output.GetRooms
}

// GetRooms calls the mock function
func (m *MockRoomUsecase) GetRooms(ctx context.Context) *output.GetRooms {
	if m.GetRoomsFunc != nil {
		return m.GetRoomsFunc(ctx)
	}
	return nil
}