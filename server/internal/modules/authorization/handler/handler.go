package handler

import (
	auth_v1 "github.com/ekkx/tcmrsv-web/server/internal/api/v1/authorization"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/usecase"
)

type Handler struct {
	auth_v1.UnimplementedAuthorizationServiceServer

	Usecase *usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}
