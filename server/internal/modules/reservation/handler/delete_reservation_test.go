package handler_test

import (
	"context"
	"testing"

	"github.com/ekkx/tcmrsv-web/server/internal/config"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/interceptor"
	mockusecase "github.com/ekkx/tcmrsv-web/server/tests/mocks/usecase"
	"github.com/ekkx/tcmrsv-web/server/tests/testhelper"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
)

func TestHandler_DeleteReservation(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	validUserID := "test-user-123"

	tests := []struct {
		name      string
		userID    string
		request   *reservation.DeleteReservationRequest
		mockSetup func(*mockusecase.MockReservationUsecase)
		checkFunc func(*testing.T, *reservation.DeleteReservationReply, error)
		withAuth  bool
	}{
		{
			name:   "正常系: 予約削除成功",
			userID: validUserID,
			request: &reservation.DeleteReservationRequest{
				ReservationId: 123,
			},
			mockSetup: func(m *mockusecase.MockReservationUsecase) {
				m.DeleteReservationFunc = func(ctx context.Context, input *input.DeleteReservation) error {
					if input.ReservationID != 123 {
						t.Fatal("unexpected reservation ID")
					}

					return nil
				}
			},
			checkFunc: func(t *testing.T, reply *reservation.DeleteReservationReply, err error) {
				require.NoError(t, err)
				require.NotNil(t, reply)
			},
			withAuth: true,
		},
		{
			name:   "異常系: 予約が存在しない",
			userID: validUserID,
			request: &reservation.DeleteReservationRequest{
				ReservationId: 999,
			},
			mockSetup: func(m *mockusecase.MockReservationUsecase) {
				m.DeleteReservationFunc = func(ctx context.Context, input *input.DeleteReservation) error {
					return errs.ErrReservationNotFound
				}
			},
			checkFunc: func(t *testing.T, reply *reservation.DeleteReservationReply, err error) {
				testhelper.AssertGRPCError(t, err, codes.NotFound, "reservation not found")
			},
			withAuth: true,
		},
		{
			name:   "異常系: 権限なし",
			userID: validUserID,
			request: &reservation.DeleteReservationRequest{
				ReservationId: 123,
			},
			mockSetup: func(m *mockusecase.MockReservationUsecase) {
				m.DeleteReservationFunc = func(ctx context.Context, input *input.DeleteReservation) error {
					return errs.ErrPermissionDenied
				}
			},
			checkFunc: func(t *testing.T, reply *reservation.DeleteReservationReply, err error) {
				testhelper.AssertGRPCError(t, err, codes.PermissionDenied, "permission denied")
			},
			withAuth: true,
		},
		{
			name:   "異常系: 認証なし",
			userID: "",
			request: &reservation.DeleteReservationRequest{
				ReservationId: 123,
			},
			checkFunc: func(t *testing.T, reply *reservation.DeleteReservationReply, err error) {
				testhelper.AssertGRPCError(t, err, codes.Unauthenticated, "")
			},
			withAuth: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックのセットアップ
			mockUsecase := &mockusecase.MockReservationUsecase{}
			if tt.mockSetup != nil {
				tt.mockSetup(mockUsecase)
			}

			// テスト用のハンドラーを作成
			testHandler := &testReservationHandler{
				mockDeleteReservation: mockUsecase.DeleteReservationFunc,
			}

			// gRPCテストサーバーのセットアップ
			authInterceptor := interceptor.AuthUnaryInterceptor(cfg.JWTSecret)
			server := testhelper.NewGRPCTestServer(authInterceptor)
			reservation.RegisterReservationServiceServer(server.GetServer(), testHandler)
			server.Start()
			defer server.Stop()

			// テストクライアントの作成
			ctx := ctxhelper.SetConfig(context.Background(), cfg)

			if tt.withAuth {
				token := testhelper.GetTestJWTToken(tt.userID, cfg.JWTSecret)
				ctx = testhelper.SetAuthorizationHeader(ctx, token)
			}

			conn := testhelper.CreateTestClient(t, ctx, server.GetDialer())
			defer conn.Close()

			client := reservation.NewReservationServiceClient(conn)

			// テスト実行
			reply, err := client.DeleteReservation(ctx, tt.request)

			// 結果の検証
			if tt.checkFunc != nil {
				tt.checkFunc(t, reply, err)
			}
		})
	}
}