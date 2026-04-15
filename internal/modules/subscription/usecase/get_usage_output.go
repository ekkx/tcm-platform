package usecase

import (
	"connectrpc.com/connect"
	subscriptionv1 "github.com/ekkx/tcm-platform/internal/gen/pb/subscription/v1"
)

type GetUsageOutput struct {
	UsedMinutes  int32
	TotalMinutes *int32
}

func (o *GetUsageOutput) ToResponse() *connect.Response[subscriptionv1.GetUsageResponse] {
	return connect.NewResponse(&subscriptionv1.GetUsageResponse{
		Usage: &subscriptionv1.Usage{
			UsedMinutes:  o.UsedMinutes,
			TotalMinutes: o.TotalMinutes,
		},
	})
}
