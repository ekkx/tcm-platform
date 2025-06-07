package handler

import (
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/usecase"
	auth_v1 "github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/authorization"
)

type Handler struct {
	auth_v1.UnimplementedAuthorizationServiceServer

	Usecase usecase.Usecase
}

func NewHandler(usecase usecase.Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}
