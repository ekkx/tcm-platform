package usecase

import (
	"context"

	"connectrpc.com/connect"
	authv1 "github.com/ekkx/tcmrsv-web/internal/shared/pb/auth/v1"
)

type ReauthorizeInput struct {
	RefreshToken string
}

func NewReauthorizeInputFromRequest(ctx context.Context, req *connect.Request[authv1.ReauthorizeRequest]) (*ReauthorizeInput, error) {
	return &ReauthorizeInput{
		RefreshToken: req.Msg.RefreshToken,
	}, nil
}

func (st *ReauthorizeInput) Validate() error {
	return nil
}
