package tcmrsv

import "github.com/ekkx/tcmrsv"

type MockTCMClient struct {
	LoginFunc            func(params *tcmrsv.LoginParams) error
	GetRoomsFilteredFunc func(params tcmrsv.GetRoomsFilteredParams) []tcmrsv.Room
	GetRoomsFunc         func() []tcmrsv.Room
}

func (m *MockTCMClient) Login(params *tcmrsv.LoginParams) error {
	return m.LoginFunc(params)
}

func (m *MockTCMClient) GetRoomsFiltered(params tcmrsv.GetRoomsFilteredParams) []tcmrsv.Room {
	return m.GetRoomsFilteredFunc(params)
}

func (m *MockTCMClient) GetRooms() []tcmrsv.Room {
	return m.GetRoomsFunc()
}
