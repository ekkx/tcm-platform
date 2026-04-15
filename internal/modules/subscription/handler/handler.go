package handler

import (
	"github.com/ekkx/tcm-platform/internal/gen/pb/subscription/v1/subscriptionv1connect"
	"github.com/ekkx/tcm-platform/internal/modules/subscription/usecase"
)

type HandlerImpl struct {
	useCase usecase.UseCase
}

func New(useCase usecase.UseCase) subscriptionv1connect.SubscriptionServiceHandler {
	return &HandlerImpl{
		useCase: useCase,
	}
}
