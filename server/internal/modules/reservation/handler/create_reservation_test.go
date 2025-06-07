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

func TestHandler_CreateReservation(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	validUserID := "test-user-123"
	validDate := time.Now().AddDate(0, 0, 1).Truncate(24 * time.Hour) // 明日

	tests := []struct {
		name          string
		userID        string
		request       *reservation.CreateReservationRequest
		mockSetup     func(*mockusecase.MockReservationUsecase)
		expectedError error
		checkFunc     func(*testing.T, *reservation.CreateReservationReply, error)
		withAuth      bool
	}{
		{
			name:   "正常系: 予約作成成功",
			userID: validUserID,
			request: &reservation.CreateReservationRequest{
				Reservation: &reservation.ReservationInput{
					CampusType: room.CampusType_IKEBUKURO,
					Date:       timestamppb.New(validDate),
					RoomId:     "room-1",
					FromHour:   10,
					FromMinute: 0,
					ToHour:     12,
					ToMinute:   0,
				},
			},
			mockSetup: func(m *mockusecase.MockReservationUsecase) {
				m.CreateReservationFunc = func(ctx context.Context, input *input.CreateReservation) (*output.CreateReservation, error) {
					// 入力値の検証
					if input.UserID != validUserID {
						t.Fatalf("expected userID %s, got %s", validUserID, input.UserID)
					}
					if input.CampusType != enum.CampusTypeIkebukuro {
						t.Fatal("unexpected campus type")
					}
					if input.RoomID != "room-1" {
						t.Fatal("unexpected room ID")
					}

					return output.NewCreateReservation([]entity.Reservation{
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
					}), nil
				}
			},
			checkFunc: func(t *testing.T, reply *reservation.CreateReservationReply, err error) {
				require.NoError(t, err)
				require.NotNil(t, reply)
				require.NotNil(t, reply.Reservations)
				require.Len(t, reply.Reservations, 1)
				require.Equal(t, int64(123), reply.Reservations[0].Id)
				require.Equal(t, "ext-123", reply.Reservations[0].GetExternalId())
				require.Equal(t, room.CampusType_IKEBUKURO, reply.Reservations[0].CampusType)
				require.Equal(t, "room-1", reply.Reservations[0].RoomId)
			},
			withAuth: true,
		},
		{
			name:   "異常系: 認証なし",
			userID: "",
			request: &reservation.CreateReservationRequest{
				Reservation: &reservation.ReservationInput{
					CampusType: room.CampusType_IKEBUKURO,
					Date:       timestamppb.New(validDate),
					RoomId:     "room-1",
					FromHour:   10,
					FromMinute: 0,
					ToHour:     12,
					ToMinute:   0,
				},
			},
			checkFunc: func(t *testing.T, reply *reservation.CreateReservationReply, err error) {
				testhelper.AssertGRPCError(t, err, codes.Unauthenticated, "")
			},
			withAuth: false,
		},
		{
			name:   "異常系: 時間の重複",
			userID: validUserID,
			request: &reservation.CreateReservationRequest{
				Reservation: &reservation.ReservationInput{
					CampusType: room.CampusType_IKEBUKURO,
					Date:       timestamppb.New(validDate),
					RoomId:     "room-1",
					FromHour:   10,
					FromMinute: 0,
					ToHour:     12,
					ToMinute:   0,
				},
			},
			mockSetup: func(m *mockusecase.MockReservationUsecase) {
				m.CreateReservationFunc = func(ctx context.Context, input *input.CreateReservation) (*output.CreateReservation, error) {
					return nil, errs.ErrReservationConflict
				}
			},
			checkFunc: func(t *testing.T, reply *reservation.CreateReservationReply, err error) {
				testhelper.AssertGRPCError(t, err, codes.AlreadyExists, "reservation conflict")
			},
			withAuth: true,
		},
		{
			name:   "異常系: 部屋が存在しない",
			userID: validUserID,
			request: &reservation.CreateReservationRequest{
				Reservation: &reservation.ReservationInput{
					CampusType: room.CampusType_IKEBUKURO,
					Date:       timestamppb.New(validDate),
					RoomId:     "non-existent-room",
					FromHour:   10,
					FromMinute: 0,
					ToHour:     12,
					ToMinute:   0,
				},
			},
			mockSetup: func(m *mockusecase.MockReservationUsecase) {
				m.CreateReservationFunc = func(ctx context.Context, input *input.CreateReservation) (*output.CreateReservation, error) {
					return nil, errs.ErrRoomNotFound
				}
			},
			checkFunc: func(t *testing.T, reply *reservation.CreateReservationReply, err error) {
				testhelper.AssertGRPCError(t, err, codes.NotFound, "room not found")
			},
			withAuth: true,
		},
		{
			name:   "異常系: 無効な時間範囲",
			userID: validUserID,
			request: &reservation.CreateReservationRequest{
				Reservation: &reservation.ReservationInput{
					CampusType: room.CampusType_IKEBUKURO,
					Date:       timestamppb.New(validDate),
					RoomId:     "room-1",
					FromHour:   14,
					FromMinute: 0,
					ToHour:     12,
					ToMinute:   0,
				},
			},
			mockSetup: func(m *mockusecase.MockReservationUsecase) {
				m.CreateReservationFunc = func(ctx context.Context, input *input.CreateReservation) (*output.CreateReservation, error) {
					return nil, errs.ErrInvalidArgument
				}
			},
			checkFunc: func(t *testing.T, reply *reservation.CreateReservationReply, err error) {
				testhelper.AssertGRPCError(t, err, codes.InvalidArgument, "")
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

			// gRPCテストサーバーのセットアップ（認証interceptor付き）
			authInterceptor := interceptor.AuthUnaryInterceptor(cfg.JWTSecret)
			server := testhelper.NewGRPCTestServer(authInterceptor)
			reservation.RegisterReservationServiceServer(server.GetServer(), testHandler)
			server.Start()
			defer server.Stop()

			// テストクライアントの作成
			ctx := ctxhelper.SetConfig(context.Background(), cfg)

			// 認証が必要な場合はトークンを設定
			if tt.withAuth {
				token := testhelper.GetTestJWTToken(tt.userID, cfg.JWTSecret)
				ctx = testhelper.SetAuthorizationHeader(ctx, token)
			}

			conn := testhelper.CreateTestClient(t, ctx, server.GetDialer())
			defer conn.Close()

			client := reservation.NewReservationServiceClient(conn)

			// テスト実行
			reply, err := client.CreateReservation(ctx, tt.request)

			// 結果の検証
			if tt.checkFunc != nil {
				tt.checkFunc(t, reply, err)
			}
		})
	}
}