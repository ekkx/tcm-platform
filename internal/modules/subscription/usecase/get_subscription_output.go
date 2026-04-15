package usecase

import (
	"connectrpc.com/connect"
	"github.com/ekkx/tcm-platform/internal/app/presenter"
	"github.com/ekkx/tcm-platform/internal/domain/entity"
	subscriptionv1 "github.com/ekkx/tcm-platform/internal/gen/pb/subscription/v1"
)

type GetSubscriptionOutput struct {
	Subscription *entity.Subscription
}

func (o *GetSubscriptionOutput) ToResponse() *connect.Response[subscriptionv1.GetSubscriptionResponse] {
	return connect.NewResponse(&subscriptionv1.GetSubscriptionResponse{
		Subscription: presenter.ToSubscription(o.Subscription),
	})
}
