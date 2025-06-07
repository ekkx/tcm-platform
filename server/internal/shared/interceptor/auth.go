package interceptor

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/jwter"
)

var publicMethods = map[string]bool{
	"/proto.v1.authorization.AuthorizationService/Authorize":   true,
	"/proto.v1.authorization.AuthorizationService/Reauthorize": true,
}

func AuthUnaryInterceptor(jwtSecret string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if isPublicMethod(info.FullMethod) {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, errs.ErrUnauthorized.Error())
		}

		authHeader := md.Get("authorization")
		if len(authHeader) == 0 {
			return nil, status.Error(codes.Unauthenticated, errs.ErrUnauthorized.Error())
		}

		token := extractToken(authHeader[0])
		if token == "" {
			return nil, status.Error(codes.Unauthenticated, errs.ErrInvalidAccessToken.Error())
		}

		userID, err := jwter.Verify(token, "access", []byte(jwtSecret))
		if err != nil {
			switch err {
			case jwter.ErrTokenExpired:
				return nil, status.Error(codes.Unauthenticated, errs.ErrAccessTokenExpired.Error())
			case jwter.ErrInvalidToken:
				return nil, status.Error(codes.Unauthenticated, errs.ErrInvalidAccessToken.Error())
			case jwter.ErrInvalidTokenScope:
				return nil, status.Error(codes.Unauthenticated, errs.ErrInvalidJWTScope.Error())
			default:
				return nil, status.Error(codes.Unauthenticated, errs.ErrInvalidAccessToken.Error())
			}
		}

		if userID == "" {
			return nil, status.Error(codes.Unauthenticated, errs.ErrInvalidAccessToken.Error())
		}

		act := actor.Actor{
			ID:   userID,
			Role: actor.RoleUser,
		}
		ctx = ctxhelper.SetActor(ctx, act)
		ctx = ctxhelper.SetAccessToken(ctx, token)

		return handler(ctx, req)
	}
}

func isPublicMethod(fullMethod string) bool {
	return publicMethods[fullMethod]
}

func extractToken(authHeader string) string {
	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return ""
	}
	return strings.TrimPrefix(authHeader, bearerPrefix)
}
