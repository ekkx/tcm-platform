package handler_test

import (
	"context"
	"testing"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/config"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/output"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/interceptor"
	mockusecase "github.com/ekkx/tcmrsv-web/server/tests/mocks/usecase"
	"github.com/ekkx/tcmrsv-web/server/tests/testhelper"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
)

func TestHandler_GetUserReservations(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	validUserID := "test-user-123"
	validDate := time.Now().AddDate(0, 0, 1).Truncate(24 * time.Hour)

	tests := []struct {
		name      string
		userID    string
		request   *reservation.GetUserReservationsRequest
		mockSetup func(*mockusecase.MockReservationUsecase)
		checkFunc func(*testing.T, *reservation.GetUserReservationsReply, error)
		withAuth  bool
	}{
		{
			name:    "正常系: ユーザーの予約一覧取得成功",
			userID:  validUserID,
			request: &reservation.GetUserReservationsRequest{},
			mockSetup: func(m *mockusecase.MockReservationUsecase) {
				m.GetUserReservationsFunc = func(ctx context.Context, input *input.GetUserReservations) (*output.GetMyReservations, error) {
					// Actor情報から取得したユーザーIDの検証
					if input.UserID != validUserID {
						t.Fatal("unexpected user ID from actor")
					}

					return output.NewGetMyReservations([]entity.Reservation{
						{
							ID:         123,
							ExternalID: ptr("ext-123"),
							UserID:     validUserID,
							CampusType: enum.CampusTypeIkebukuro,
							Date:       validDate,
							RoomID:     "room-1",
							FromHour:   10,
							FromMinute: 0,
							ToHour:     12,
							ToMinute:   0,
							BookerName: ptr("Test User"),
							CreatedAt:  time.Now(),
						},
						{
							ID:         124,
							ExternalID: ptr("ext-124"),
							UserID:     validUserID,
							CampusType: enum.CampusTypeNakameguro,
							Date:       validDate.AddDate(0, 0, 1),
							RoomID:     "room-2",
							FromHour:   14,
							FromMinute: 30,
							ToHour:     16,
							ToMinute:   0,
							BookerName: ptr("Test User"),
							CreatedAt:  time.Now(),
						},
					}), nil
				}
			},
			checkFunc: func(t *testing.T, reply *reservation.GetUserReservationsReply, err error) {
				require.NoError(t, err)
				require.NotNil(t, reply)
				require.Len(t, reply.Reservations, 2)
				require.Equal(t, int64(123), reply.Reservations[0].Id)
				require.Equal(t, int64(124), reply.Reservations[1].Id)
			},
			withAuth: true,
		},
		{
			name:    "正常系: 予約がない場合",
			userID:  validUserID,
			request: &reservation.GetUserReservationsRequest{},
			mockSetup: func(m *mockusecase.MockReservationUsecase) {
				m.GetUserReservationsFunc = func(ctx context.Context, input *input.GetUserReservations) (*output.GetMyReservations, error) {
					return output.NewGetMyReservations([]entity.Reservation{}), nil
				}
			},
			checkFunc: func(t *testing.T, reply *reservation.GetUserReservationsReply, err error) {
				require.NoError(t, err)
				require.NotNil(t, reply)
				require.Empty(t, reply.Reservations)
			},
			withAuth: true,
		},
		{
			name:    "異常系: 認証なし",
			userID:  "",
			request: &reservation.GetUserReservationsRequest{},
			checkFunc: func(t *testing.T, reply *reservation.GetUserReservationsReply, err error) {
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
				mockCreateReservation:   mockUsecase.CreateReservationFunc,
				mockGetReservation:      mockUsecase.GetReservationFunc,
				mockGetUserReservations: mockUsecase.GetUserReservationsFunc,
				mockUpdateReservation:   mockUsecase.UpdateReservationFunc,
				mockDeleteReservation:   mockUsecase.DeleteReservationFunc,
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
			reply, err := client.GetMyReservations(ctx, tt.request)

			// 結果の検証
			if tt.checkFunc != nil {
				tt.checkFunc(t, reply, err)
			}
		})
	}
}