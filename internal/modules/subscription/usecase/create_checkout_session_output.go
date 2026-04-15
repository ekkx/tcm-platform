package usecase

import (
	"connectrpc.com/connect"
	subscriptionv1 "github.com/ekkx/tcm-platform/internal/gen/pb/subscription/v1"
)

type CreateCheckoutSessionOutput struct {
	CheckoutURL string
}

func (o *CreateCheckoutSessionOutput) ToResponse() *connect.Response[subscriptionv1.CreateCheckoutSessionResponse] {
	return connect.NewResponse(&subscriptionv1.CreateCheckoutSessionResponse{
		CheckoutUrl: o.CheckoutURL,
	})
}
