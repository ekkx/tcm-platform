package usecase_test

import (
	"testing"
	"time"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	rsv_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	rsv_uc "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/usecase"
	room_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/room/repository"
	user_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/ekkx/tcmrsv-web/server/pkg/utils"
	"github.com/ekkx/tcmrsv-web/server/tests/testhelper"
	"github.com/stretchr/testify/require"
)

func TestCreateReservation_正常系(t *testing.T) {
	t.Run("練習室を指定して予約作成", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			roomRepo := room_repo.NewRepository(tcmrsv.New())
			rsvRepo := rsv_repo.NewRepository(db)
			userRepo := user_repo.NewRepository(db)
			rsvUC := rsv_uc.NewUsecase(rsvRepo, roomRepo)

			// ユーザーを作成
			_, err := userRepo.CreateUser(ctx, &user_repo.CreateUserArgs{
				ID:                "testuser",
				EncryptedPassword: "testpass",
			})
			require.NoError(t, err)

			// ルームを取得
			rooms := roomRepo.SearchRooms(ctx, &room_repo.SearchRoomsArgs{})
			require.NotEmpty(t, rooms)

			// room.CampusType が enum.CampusTypeIkebukuro の練習室を一つ取得
			var room entity.Room
			for _, r := range rooms {
				if r.CampusType == enum.CampusTypeIkebukuro {
					room = r
					break
				}
			}

			// 予約を作成
			date := time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST())
			output, err := rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       date,
				FromHour:   10,
				FromMinute: 0,
				ToHour:     12,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: nil,
			})
			require.NoError(t, err)

			// 予約の内容を確認
			require.Len(t, output.Reservations, 1)
			require.Equal(t, room.ID, output.Reservations[0].RoomID)
			require.Equal(t, enum.CampusTypeIkebukuro, output.Reservations[0].CampusType)
			require.Equal(t, time.Date(2033, 10, 1, 0, 0, 0, 0, utils.JST()), output.Reservations[0].Date)
			require.Equal(t, int32(10), output.Reservations[0].FromHour)
			require.Equal(t, int32(0), output.Reservations[0].FromMinute)
			require.Equal(t, int32(12), output.Reservations[0].ToHour)
			require.Equal(t, int32(0), output.Reservations[0].ToMinute)
		})
	})
}
