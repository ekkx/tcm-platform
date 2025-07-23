package interceptor

import (
	"context"
	"strings"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/actor"
	"github.com/ekkx/tcmrsv-web/server/pkg/jwt"
	"github.com/ekkx/tcmrsv-web/server/pkg/ulid"
)

const bearerPrefix = "Bearer "

// AuthInterceptor は JWT を使用してリクエストの認証を行うインターセプター。
// トークンの検証のみを行うので、ユーザーの存在チェックや権限チェックは行わない。
func AuthInterceptor(jwtManager *jwt.JWTManager) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			authHeader := req.Header().Get("Authorization")
			if !strings.HasPrefix(authHeader, bearerPrefix) {
				return nil, errs.ErrInvalidAuthorizationHeader
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")

			claims, err := jwtManager.VerifyToken(token)
			if err != nil {
				switch err {
				case jwt.ErrInvalidToken:
					return nil, errs.ErrInvalidToken
				case jwt.ErrExpiredToken:
					return nil, errs.ErrExpiredToken
				default:
					return nil, errs.ErrInternal.WithCause(err)
				}
			}

			// アクセストークンのみを許可する
			if claims.TokenType != jwt.TokenTypeAccess {
				return nil, errs.ErrInvalidTokenType
			}

			userID, err := ulid.Parse(claims.Subject)
			if err != nil {
				// サーバー側で生成したトークンなので ID が不正なのは、サーバー内部の構造的ミス
				return nil, errs.ErrInternal.WithCause(err)
			}

			actor := actor.New(userID, actor.TypeUser)

			newCtx := ctxhelper.WithActor(ctx, actor)
			return next(newCtx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
