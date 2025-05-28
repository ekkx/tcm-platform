package repository

import (
	"github.com/ekkx/tcmrsv-web/server/internal/domain/port"
)

type Repository struct {
	tcmClient port.TCMClient
}

func NewRepository(tcmClient port.TCMClient) *Repository {
	return &Repository{
		tcmClient: tcmClient,
	}
}
