package usecase

import (
	"connectrpc.com/connect"
	subscriptionv1 "github.com/ekkx/tcm-platform/internal/gen/pb/subscription/v1"
)

type CreatePortalSessionOutput struct {
	PortalURL string
}

func (o *CreatePortalSessionOutput) ToResponse() *connect.Response[subscriptionv1.CreatePortalSessionResponse] {
	return connect.NewResponse(&subscriptionv1.CreatePortalSessionResponse{
		PortalUrl: o.PortalURL,
	})
}
