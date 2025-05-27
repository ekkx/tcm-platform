package tcmrsv

import "github.com/ekkx/tcmrsv"

type MockTCMClient struct {
	LoginFunc func(params *tcmrsv.LoginParams) error
}

func (m *MockTCMClient) Login(params *tcmrsv.LoginParams) error {
	return m.LoginFunc(params)
}
