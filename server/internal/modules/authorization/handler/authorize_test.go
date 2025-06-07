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

func TestHandler_Authorize(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	tests := []struct {
		name           string
		request        *authorization.AuthorizeRequest
		mockSetup      func(*usecase.MockAuthorizationUsecase)
		expectedError  error
		expectedOutput *authorization.AuthorizeReply
		checkFunc      func(*testing.T, *authorization.AuthorizeReply, error)
	}{
		{
			name: "正常系: 認証成功",
			request: &authorization.AuthorizeRequest{
				UserId:   "test-user",
				Password: "test-password",
			},
			mockSetup: func(m *usecase.MockAuthorizationUsecase) {
				m.AuthorizeFunc = func(ctx context.Context, input *input.Authorize) (*output.Authorize, error) {
					// 入力値の検証
					if input.UserID != "test-user" || input.Password != "test-password" {
						t.Fatal("unexpected input")
					}
					// Configから値が設定されているか確認
					if input.JWTSecret == "" || input.PasswordAESKey == "" {
						t.Fatal("config values not set")
					}
					
					return output.NewAuthorize(entity.Authorization{
						AccessToken:  "test-access-token",
						RefreshToken: "test-refresh-token",
					}), nil
				}
			},
			expectedOutput: &authorization.AuthorizeReply{
				Authorization: &authorization.Authorization{
					AccessToken:  "test-access-token",
					RefreshToken: "test-refresh-token",
				},
			},
		},
		{
			name: "異常系: 認証失敗 - ユーザーが存在しない",
			request: &authorization.AuthorizeRequest{
				UserId:   "non-existent-user",
				Password: "test-password",
			},
			mockSetup: func(m *usecase.MockAuthorizationUsecase) {
				m.AuthorizeFunc = func(ctx context.Context, input *input.Authorize) (*output.Authorize, error) {
					return nil, errs.ErrUserNotFound
				}
			},
			checkFunc: func(t *testing.T, reply *authorization.AuthorizeReply, err error) {
				testhelper.AssertGRPCError(t, err, codes.NotFound, "user not found")
			},
		},
		{
			name: "異常系: 認証失敗 - パスワードが間違っている",
			request: &authorization.AuthorizeRequest{
				UserId:   "test-user",
				Password: "wrong-password",
			},
			mockSetup: func(m *usecase.MockAuthorizationUsecase) {
				m.AuthorizeFunc = func(ctx context.Context, input *input.Authorize) (*output.Authorize, error) {
					return nil, errs.ErrInvalidEmailOrPassword
				}
			},
			checkFunc: func(t *testing.T, reply *authorization.AuthorizeReply, err error) {
				testhelper.AssertGRPCError(t, err, codes.InvalidArgument, "invalid email or password")
			},
		},
		{
			name: "異常系: 空のユーザーID",
			request: &authorization.AuthorizeRequest{
				UserId:   "",
				Password: "test-password",
			},
			mockSetup: func(m *usecase.MockAuthorizationUsecase) {
				m.AuthorizeFunc = func(ctx context.Context, input *input.Authorize) (*output.Authorize, error) {
					return nil, errs.ErrInvalidArgument
				}
			},
			checkFunc: func(t *testing.T, reply *authorization.AuthorizeReply, err error) {
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
			reply, err := client.Authorize(ctx, tt.request)

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