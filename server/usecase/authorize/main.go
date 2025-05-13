package authorize

import (
	"context"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/infra/db"
)

type AuthorizeUsecase interface {
	Authorize(ctx context.Context, input *AuthorizeInput) (*AuthorizeOutput, error)
	Reauthorize(ctx context.Context, input *ReauthorizeInput) (*AuthorizeOutput, error)
}

type AuthorizeUsecaseImpl struct {
	tcmClient *tcmrsv.Client
	querier   db.Querier
}

var _ AuthorizeUsecase = (*AuthorizeUsecaseImpl)(nil)

func New(tcmClient *tcmrsv.Client, querier db.Querier) AuthorizeUsecase {
	return &AuthorizeUsecaseImpl{
		tcmClient: tcmClient,
		querier:   querier,
	}
}
