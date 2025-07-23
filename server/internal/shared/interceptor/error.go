package interceptor

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/config"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
)

func ErrorInterceptor(env config.Env) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			resp, err := next(ctx, req)
			if err == nil {
				return resp, nil
			}

			var appErr *errs.Error
			if errors.As(err, &appErr) {
				return nil, connect.NewError(appErr.ConnectCode, appErr)
			}

			// 開発環境でのみ詳細なエラーメッセージを返す
			if env == config.EnvDevelopment {
				return nil, connect.NewError(connect.CodeInternal, err)
			}

			return nil, connect.NewError(connect.CodeInternal, errors.New("internal server error"))
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
