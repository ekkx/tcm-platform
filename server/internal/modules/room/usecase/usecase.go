package usecase

import "github.com/ekkx/tcmrsv-web/server/internal/core/port"

type Usecase struct {
	tcmClient port.TCMClient
}

func NewUsecase(tcmClient port.TCMClient) *Usecase {
	return &Usecase{
		tcmClient: tcmClient,
	}
}
