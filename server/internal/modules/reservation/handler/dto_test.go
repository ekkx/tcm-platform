package handler_test

import (
	"context"
	"testing"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/config"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/output"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/reservation"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/room"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	mockusecase "github.com/ekkx/tcmrsv-web/server/tests/mocks/usecase"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TestHandler_DTOConversion tests DTO conversion logic in handler
func TestHandler_DTOConversion(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	t.Run("input DTO がコンテキストからActor情報を正しく取得すること", func(t *testing.T) {
		validUserID := "test-user-123"

		mockUsecase := &mockusecase.MockReservationUsecase{
			CreateReservationFunc: func(ctx context.Context, input *input.CreateReservation) (*output.CreateReservation, error) {
				// InputDTOがコンテキストからActor情報を正しく取得しているか確認
				require.Equal(t, validUserID, input.UserID)

				return output.NewCreateReservation([]entity.Reservation{
					{
						ID:     123,
						UserID: validUserID,
					},
				}), nil
			},
		}

		// テスト用のハンドラーを作成
		testHandler := &testReservationHandler{
			mockCreateReservation: mockUsecase.CreateReservationFunc,
		}

		// コンテキストにActor情報を設定
		ctx := ctxhelper.SetConfig(context.Background(), cfg)
		ctx = ctxhelper.SetActor(ctx, actor.Actor{
			ID: validUserID,
		})

		req := &reservation.CreateReservationRequest{
			Reservation: &reservation.ReservationInput{
				CampusType: room.CampusType_IKEBUKURO,
				Date:       timestamppb.New(time.Now().AddDate(0, 0, 1)),
				RoomId:     "room-1",
				FromHour:   10,
				FromMinute: 0,
				ToHour:     12,
				ToMinute:   0,
			},
		}

		reply, err := testHandler.CreateReservation(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, reply)
	})

	t.Run("output DTO が空の予約リストを正しく処理すること", func(t *testing.T) {
		mockUsecase := &mockusecase.MockReservationUsecase{
			GetUserReservationsFunc: func(ctx context.Context, input *input.GetUserReservations) (*output.GetMyReservations, error) {
				// 空のリストを返す
				return output.NewGetMyReservations([]entity.Reservation{}), nil
			},
		}

		// テスト用のハンドラーを作成
		testHandler := &testReservationHandler{
			mockGetUserReservations: mockUsecase.GetUserReservationsFunc,
		}

		ctx := ctxhelper.SetConfig(context.Background(), cfg)
		ctx = ctxhelper.SetActor(ctx, actor.Actor{
			ID: "test-user",
		})

		req := &reservation.GetUserReservationsRequest{}

		reply, err := testHandler.GetUserReservations(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, reply)
		require.NotNil(t, reply.Reservations)
		require.Empty(t, reply.Reservations)
	})
}
