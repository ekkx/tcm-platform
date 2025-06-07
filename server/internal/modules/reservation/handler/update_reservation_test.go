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
	"github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/room"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/interceptor"
	mockusecase "github.com/ekkx/tcmrsv-web/server/tests/mocks/usecase"
	"github.com/ekkx/tcmrsv-web/server/tests/testhelper"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestHandler_UpdateReservation(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	validUserID := "test-user-123"
	validDate := time.Now().AddDate(0, 0, 1).Truncate(24 * time.Hour)

	tests := []struct {
		name      string
		userID    string
		request   *reservation.UpdateReservationRequest
		mockSetup func(*mockusecase.MockReservationUsecase)
		checkFunc func(*testing.T, *reservation.UpdateReservationReply, error)
		withAuth  bool
	}{
		{
			name:   "正常系: 予約更新成功",
			userID: validUserID,
			request: &reservation.UpdateReservationRequest{
				ReservationId: 123,
				Reservation: &reservation.ReservationInput{
					CampusType: room.CampusType_IKEBUKURO,
					Date:       timestamppb.New(validDate),
					RoomId:     "room-1",
					FromHour:   14,
					FromMinute: 0,
					ToHour:     16,
					ToMinute:   0,
				},
			},
			mockSetup: func(m *mockusecase.MockReservationUsecase) {
				m.UpdateReservationFunc = func(ctx context.Context, input *input.UpdateReservation) (*output.UpdateReservation, error) {
					if input.ID != 123 {
						t.Fatal("unexpected reservation ID")
					}

					return output.NewUpdateReservation(entity.Reservation{
						ID:         123,
						ExternalID: ptr("ext-123"),
						UserID:     validUserID,
						CampusType: enum.CampusTypeIkebukuro,
						Date:       validDate,
						RoomID:     "room-1",
						FromHour:   14,
						FromMinute: 0,
						ToHour:     16,
						ToMinute:   0,
						BookerName: ptr("Test User"),
						CreatedAt:  time.Now(),
					}), nil
				}
			},
			checkFunc: func(t *testing.T, reply *reservation.UpdateReservationReply, err error) {
				require.NoError(t, err)
				require.NotNil(t, reply)
				require.NotNil(t, reply.Reservation)
				require.Equal(t, int32(14), reply.Reservation.FromHour)
				require.Equal(t, int32(16), reply.Reservation.ToHour)
			},
			withAuth: true,
		},
		{
			name:   "異常系: 予約が存在しない",
			userID: validUserID,
			request: &reservation.UpdateReservationRequest{
				ReservationId: 999,
				Reservation: &reservation.ReservationInput{
					CampusType: room.CampusType_IKEBUKURO,
					Date:       timestamppb.New(validDate),
					RoomId:     "room-1",
					FromHour:   14,
					FromMinute: 0,
					ToHour:     16,
					ToMinute:   0,
				},
			},
			mockSetup: func(m *mockusecase.MockReservationUsecase) {
				m.UpdateReservationFunc = func(ctx context.Context, input *input.UpdateReservation) (*output.UpdateReservation, error) {
					return nil, errs.ErrReservationNotFound
				}
			},
			checkFunc: func(t *testing.T, reply *reservation.UpdateReservationReply, err error) {
				testhelper.AssertGRPCError(t, err, codes.NotFound, "reservation not found")
			},
			withAuth: true,
		},
		{
			name:   "異常系: 権限なし",
			userID: validUserID,
			request: &reservation.UpdateReservationRequest{
				ReservationId: 123,
				Reservation: &reservation.ReservationInput{
					CampusType: room.CampusType_IKEBUKURO,
					Date:       timestamppb.New(validDate),
					RoomId:     "room-1",
					FromHour:   14,
					FromMinute: 0,
					ToHour:     16,
					ToMinute:   0,
				},
			},
			mockSetup: func(m *mockusecase.MockReservationUsecase) {
				m.UpdateReservationFunc = func(ctx context.Context, input *input.UpdateReservation) (*output.UpdateReservation, error) {
					return nil, errs.ErrPermissionDenied
				}
			},
			checkFunc: func(t *testing.T, reply *reservation.UpdateReservationReply, err error) {
				testhelper.AssertGRPCError(t, err, codes.PermissionDenied, "permission denied")
			},
			withAuth: true,
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
			reply, err := client.UpdateReservation(ctx, tt.request)

			// 結果の検証
			if tt.checkFunc != nil {
				tt.checkFunc(t, reply, err)
			}
		})
	}
}