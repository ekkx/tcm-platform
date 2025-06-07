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
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/tests/mocks/usecase"
	"github.com/ekkx/tcmrsv-web/server/tests/testhelper"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
)

func TestHandler_Reauthorize(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	tests := []struct {
		name           string
		request        *authorization.ReauthorizeRequest
		mockSetup      func(*usecase.MockAuthorizationUsecase)
		expectedError  error
		expectedOutput *authorization.ReauthorizeReply
		checkFunc      func(*testing.T, *authorization.ReauthorizeReply, error)
	}{
		{
			name: "正常系: 再認証成功",
			request: &authorization.ReauthorizeRequest{
				RefreshToken: "valid-refresh-token",
			},
			mockSetup: func(m *usecase.MockAuthorizationUsecase) {
				m.ReauthorizeFunc = func(ctx context.Context, input *input.Reauthorize) (*output.Reauthorize, error) {
					// 入力値の検証
					if input.RefreshToken != "valid-refresh-token" {
						t.Fatal("unexpected refresh token")
					}
					// Configから値が設定されているか確認
					if input.JWTSecret == "" {
						t.Fatal("jwt secret not set from config")
					}
					
					return output.NewReauthorize(entity.Authorization{
						AccessToken:  "new-access-token",
						RefreshToken: "new-refresh-token",
					}), nil
				}
			},
			expectedOutput: &authorization.ReauthorizeReply{
				Authorization: &authorization.Authorization{
					AccessToken:  "new-access-token",
					RefreshToken: "new-refresh-token",
				},
			},
		},
		{
			name: "異常系: 無効なリフレッシュトークン",
			request: &authorization.ReauthorizeRequest{
				RefreshToken: "invalid-refresh-token",
			},
			mockSetup: func(m *usecase.MockAuthorizationUsecase) {
				m.ReauthorizeFunc = func(ctx context.Context, input *input.Reauthorize) (*output.Reauthorize, error) {
					return nil, errs.ErrInvalidRefreshToken
				}
			},
			checkFunc: func(t *testing.T, reply *authorization.ReauthorizeReply, err error) {
				testhelper.AssertGRPCError(t, err, codes.Unauthenticated, "invalid refresh token")
			},
		},
		{
			name: "異常系: 期限切れのリフレッシュトークン",
			request: &authorization.ReauthorizeRequest{
				RefreshToken: "expired-refresh-token",
			},
			mockSetup: func(m *usecase.MockAuthorizationUsecase) {
				m.ReauthorizeFunc = func(ctx context.Context, input *input.Reauthorize) (*output.Reauthorize, error) {
					return nil, errs.ErrRefreshTokenExpired
				}
			},
			checkFunc: func(t *testing.T, reply *authorization.ReauthorizeReply, err error) {
				testhelper.AssertGRPCError(t, err, codes.Unauthenticated, "refresh token expired")
			},
		},
		{
			name: "異常系: 空のリフレッシュトークン",
			request: &authorization.ReauthorizeRequest{
				RefreshToken: "",
			},
			mockSetup: func(m *usecase.MockAuthorizationUsecase) {
				m.ReauthorizeFunc = func(ctx context.Context, input *input.Reauthorize) (*output.Reauthorize, error) {
					return nil, errs.ErrInvalidArgument
				}
			},
			checkFunc: func(t *testing.T, reply *authorization.ReauthorizeReply, err error) {
				testhelper.AssertGRPCError(t, err, codes.InvalidArgument, "")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックのセットアップ
			mockUsecase := &usecase.MockAuthorizationUsecase{}
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// ハンドラーの作成
			h := handler.NewHandler(mockUsecase)

			// gRPCテストサーバーのセットアップ
			server := testhelper.NewGRPCTestServer()
			authorization.RegisterAuthorizationServiceServer(server.GetServer(), h)
			server.Start()
			defer server.Stop()

			// テストクライアントの作成
			ctx := ctxhelper.SetConfig(context.Background(), cfg)
			conn := testhelper.CreateTestClient(t, ctx, server.GetDialer())
			defer conn.Close()

			client := authorization.NewAuthorizationServiceClient(conn)

			// テスト実行
			reply, err := client.Reauthorize(ctx, tt.request)

			// 結果の検証
			if tt.checkFunc != nil {
				tt.checkFunc(t, reply, err)
			} else if tt.expectedError != nil {
				require.ErrorIs(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedOutput.Authorization.AccessToken, reply.Authorization.AccessToken)
				require.Equal(t, tt.expectedOutput.Authorization.RefreshToken, reply.Authorization.RefreshToken)
			}
		})
	}
}