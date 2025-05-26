package handler

import (
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorize/usecase"
	authorize_v1 "github.com/ekkx/tcmrsv-web/server/pkg/api/v1/authorize"
)

type Handler struct {
	authorize_v1.UnimplementedAuthorizeServiceServer

	Usecase *usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}
