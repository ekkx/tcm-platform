package handler_test

import (
	"context"
	"testing"

	"github.com/ekkx/tcmrsv-web/server/internal/config"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/dto/output"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/authorization/handler"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/authorization"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/interceptor"
	"github.com/ekkx/tcmrsv-web/server/tests/mocks/usecase"
	"github.com/ekkx/tcmrsv-web/server/tests/testhelper"
	"github.com/stretchr/testify/require"
)

// TestHandler_WithAuthInterceptor tests the handler with auth interceptor
func TestHandler_WithAuthInterceptor(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	// モックのセットアップ
	mockUsecase := &usecase.MockAuthorizationUsecase{
		AuthorizeFunc: func(ctx context.Context, input *input.Authorize) (*output.Authorize, error) {
			return output.NewAuthorize(entity.Authorization{
				AccessToken:  "test-access-token",
				RefreshToken: "test-refresh-token",
			}), nil
		},
		ReauthorizeFunc: func(ctx context.Context, input *input.Reauthorize) (*output.Reauthorize, error) {
			return output.NewReauthorize(entity.Authorization{
				AccessToken:  "new-access-token",
				RefreshToken: "new-refresh-token",
			}), nil
		},
	}

	// ハンドラーの作成
	h := handler.NewHandler(mockUsecase)

	// Auth Interceptorを含むgRPCテストサーバーのセットアップ
	authInterceptor := interceptor.AuthUnaryInterceptor(cfg.JWTSecret)
	server := testhelper.NewGRPCTestServer(authInterceptor)
	authorization.RegisterAuthorizationServiceServer(server.GetServer(), h)
	server.Start()
	defer server.Stop()

	// テストクライアントの作成
	ctx := ctxhelper.SetConfig(context.Background(), cfg)
	conn := testhelper.CreateTestClient(t, ctx, server.GetDialer())
	defer conn.Close()

	client := authorization.NewAuthorizationServiceClient(conn)

	t.Run("認証不要なエンドポイント（Authorize）はトークンなしでアクセス可能", func(t *testing.T) {
		req := &authorization.AuthorizeRequest{
			UserId:   "test-user",
			Password: "test-password",
		}

		reply, err := client.Authorize(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, reply)
		require.Equal(t, "test-access-token", reply.Authorization.AccessToken)
	})

	t.Run("認証不要なエンドポイント（Reauthorize）はトークンなしでアクセス可能", func(t *testing.T) {
		req := &authorization.ReauthorizeRequest{
			RefreshToken: "valid-refresh-token",
		}

		reply, err := client.Reauthorize(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, reply)
		require.Equal(t, "new-access-token", reply.Authorization.AccessToken)
	})

	t.Run("認証不要なエンドポイントでもトークンがあれば検証はされる", func(t *testing.T) {
		// 無効なトークンを設定
		ctxWithInvalidToken := testhelper.SetAuthorizationHeader(ctx, "invalid-token")
		
		req := &authorization.AuthorizeRequest{
			UserId:   "test-user",
			Password: "test-password",
		}

		// 認証不要なエンドポイントなので、無効なトークンがあってもアクセス可能
		reply, err := client.Authorize(ctxWithInvalidToken, req)
		require.NoError(t, err)
		require.NotNil(t, reply)
	})
}