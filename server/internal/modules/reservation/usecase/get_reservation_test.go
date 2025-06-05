package usecase_test

import (
	"testing"
	"time"

	"github.com/ekkx/tcmrsv"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/ekkx/tcmrsv-web/server/pkg/utils"
	"github.com/ekkx/tcmrsv-web/server/tests/testhelper"
	"github.com/stretchr/testify/require"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	rsv_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/repository"
	rsv_uc "github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/usecase"
	room_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/room/repository"
	user_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
)

func TestGetReservation_正常系(t *testing.T) {
	// TODO: tcmrsv.GetRoomsFiltered のモックを作る

	t.Run("自分の予約取得", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			rsvRepo := rsv_repo.NewRepository(db)
			userRepo := user_repo.NewRepository(db)
			roomRepo := room_repo.NewRepository(tcmrsv.New())
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
			bookerName := "Test Booker"
			output, err := rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
				FromHour:   9,
				FromMinute: 30,
				ToHour:     12,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: &bookerName,
			})
			require.NoError(t, err)
			require.Len(t, output.Reservations, 1)
			require.NotEmpty(t, output.Reservations[0].ID)
			require.Nil(t, output.Reservations[0].ExternalID)
			require.Equal(t, "testuser", output.Reservations[0].UserID)
			require.Equal(t, enum.CampusTypeIkebukuro, output.Reservations[0].CampusType)
			require.Equal(t, room.ID, output.Reservations[0].RoomID)
			require.Equal(t, time.Date(2033, 10, 1, 0, 0, 0, 0, utils.JST()), output.Reservations[0].Date)
			require.Equal(t, 9, output.Reservations[0].FromHour)
			require.Equal(t, 30, output.Reservations[0].FromMinute)
			require.Equal(t, 12, output.Reservations[0].ToHour)
			require.Equal(t, 0, output.Reservations[0].ToMinute)
			require.Equal(t, bookerName, *output.Reservations[0].BookerName)
			require.NotEmpty(t, output.Reservations[0].CreatedAt)
		})
	})

	t.Run("システムがユーザーの予約を取得する", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			rsvRepo := rsv_repo.NewRepository(db)
			userRepo := user_repo.NewRepository(db)
			roomRepo := room_repo.NewRepository(tcmrsv.New())
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
			bookerName := "Test Booker"
			output, err := rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
				FromHour:   9,
				FromMinute: 30,
				ToHour:     12,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: &bookerName,
			})
			require.NoError(t, err)
			require.Len(t, output.Reservations, 1)

			// システムが予約を取得
			output2, err := rsvUC.GetReservation(ctx, &input.GetReservation{
				Actor: actor.Actor{
					ID:   "system",
					Role: actor.RoleSystem,
				},
				ReservationID: output.Reservations[0].ID,
			})
			require.NoError(t, err)
			require.NotNil(t, output2)
			require.Equal(t, output.Reservations[0].ID, output2.Reservation.ID)
		})
	})
}

func TestGetReservation_異常系(t *testing.T) {
	// TODO: tcmrsv.GetRoomsFiltered のモックを作る

	t.Run("予約が存在しない", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			rsvRepo := rsv_repo.NewRepository(db)
			roomRepo := room_repo.NewRepository(tcmrsv.New())
			rsvUC := rsv_uc.NewUsecase(rsvRepo, roomRepo)

			// 存在しない予約IDで取得を試みる
			_, err := rsvUC.GetReservation(ctx, &input.GetReservation{
				Actor: actor.Actor{
					ID:   "testuser",
					Role: actor.RoleUser,
				},
				ReservationID: 9999, // 存在しないID
			})
			require.Error(t, err)
			require.ErrorIs(t, err, errs.ErrReservationNotFound)
		})
	})

	t.Run("ユーザーが他人の予約を取得しようとする", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			userRepo := user_repo.NewRepository(db)
			rsvRepo := rsv_repo.NewRepository(db)
			roomRepo := room_repo.NewRepository(tcmrsv.New())
			rsvUC := rsv_uc.NewUsecase(rsvRepo, roomRepo)

			// ユーザーを作成
			_, err := userRepo.CreateUser(ctx, &user_repo.CreateUserArgs{
				ID:                "testuser",
				EncryptedPassword: "testpass",
			})
			require.NoError(t, err)

			// 他のユーザーを作成
			_, err = userRepo.CreateUser(ctx, &user_repo.CreateUserArgs{
				ID:                "otheruser",
				EncryptedPassword: "otherpass",
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

			// 他人の予約を作成
			bookerName := "Test Booker"
			output, err := rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "otheruser",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
				FromHour:   9,
				FromMinute: 30,
				ToHour:     12,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: &bookerName,
			})
			require.NoError(t, err)
			require.Len(t, output.Reservations, 1)

			// 他人の予約を取得しようとする
			_, err = rsvUC.GetReservation(ctx, &input.GetReservation{
				Actor: actor.Actor{
					ID:   "testuser",
					Role: actor.RoleUser,
				},
				ReservationID: output.Reservations[0].ID,
			})
			require.Error(t, err)
			require.ErrorIs(t, err, errs.ErrNotYourReservation)
		})
	})
}
