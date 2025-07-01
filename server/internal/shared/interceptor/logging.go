package interceptor

import (
	"context"
	"log/slog"
	"time"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
)

func NewLoggingInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			requestID := ulid.New().String()

			logger := slog.With(slog.String("request_id", requestID))
			newCtx := ctxhelper.WithLogger(ctx, logger)

			// リクエスト情報取得
			procedure := req.Spec().Procedure
			peer := req.Peer()

			logger.Info(
				"[REQ]",
				slog.String("procedure", procedure),
				slog.String("protocol", peer.Protocol),
				slog.String("peer_addr", peer.Addr),
			)

			start := time.Now()

			resp, err := next(newCtx, req)
			duration := time.Since(start)
			if err != nil {
				logger.Error(
					"[RES]",
					slog.String("procedure", procedure),
					slog.String("protocol", peer.Protocol),
					slog.String("peer_addr", peer.Addr),
					slog.Duration("duration", duration),
					slog.Any("err", err),
				)
				return resp, err
			}

			logger.Info(
				"[RES]",
				slog.String("procedure", procedure),
				slog.String("protocol", peer.Protocol),
				slog.String("peer_addr", peer.Addr),
				slog.Duration("duration", duration),
			)

			return resp, nil
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
