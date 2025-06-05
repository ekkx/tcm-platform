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
	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/ekkx/tcmrsv-web/server/pkg/utils"
	"github.com/ekkx/tcmrsv-web/server/tests/testhelper"
	"github.com/stretchr/testify/require"
)

func TestDeleteReservation_正常系(t *testing.T) {
	// TODO: tcmrsv.GetRoomsFiltered のモックを作る

	t.Run("予約削除", func(t *testing.T) {
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

			// 予約が作成されたことを確認
			_, err = rsvRepo.GetReservationByID(ctx, output.Reservations[0].ID)
			require.NoError(t, err)

			// 予約を削除
			err = rsvUC.DeleteReservation(ctx, &input.DeleteReservation{
				Actor: actor.Actor{
					ID:   "testuser",
					Role: actor.RoleUser,
				},
				ReservationID: output.Reservations[0].ID,
			})
			require.NoError(t, err)

			// 予約が削除されたことを確認
			_, err = rsvRepo.GetReservationByID(ctx, output.Reservations[0].ID)
			require.ErrorIs(t, err, errs.ErrReservationNotFound)
		})
	})

	t.Run("システムユーザーによる予約削除", func(t *testing.T) {
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

			// testuser の予約を作成
			output, err := rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
				FromHour:   9,
				FromMinute: 30,
				ToHour:     12,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: nil,
			})
			require.NoError(t, err)
			require.Len(t, output.Reservations, 1)

			// システムが予約を削除
			err = rsvUC.DeleteReservation(ctx, &input.DeleteReservation{
				Actor: actor.Actor{
					ID:   "testsystem",
					Role: actor.RoleSystem,
				},
				ReservationID: output.Reservations[0].ID,
			})
			require.NoError(t, err)

			// 予約が削除されたことを確認
			_, err = rsvRepo.GetReservationByID(ctx, output.Reservations[0].ID)
			require.ErrorIs(t, err, errs.ErrReservationNotFound)
		})
	})
}

func TestDeleteReservation_異常系(t *testing.T) {
	// TODO: tcmrsv.GetRoomsFiltered のモックを作る

	t.Run("存在しない予約の削除", func(t *testing.T) {
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

			nonExistentReservationID := 99999 // 存在しない予約ID

			// 予約が存在しないことを確認
			_, err = rsvRepo.GetReservationByID(ctx, nonExistentReservationID)
			require.ErrorIs(t, err, errs.ErrReservationNotFound)

			// 予約を削除 (user として)
			err = rsvUC.DeleteReservation(ctx, &input.DeleteReservation{
				Actor: actor.Actor{
					ID:   "testuser",
					Role: actor.RoleUser,
				},
				ReservationID: nonExistentReservationID,
			})
			require.ErrorIs(t, err, errs.ErrReservationNotFound)

			// 予約を削除 (system として)
			err = rsvUC.DeleteReservation(ctx, &input.DeleteReservation{
				Actor: actor.Actor{
					ID:   "testsystem",
					Role: actor.RoleSystem,
				},
				ReservationID: nonExistentReservationID,
			})
			require.ErrorIs(t, err, errs.ErrReservationNotFound)
		})
	})

	t.Run("他のユーザーの予約を削除", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			roomRepo := room_repo.NewRepository(tcmrsv.New())
			rsvRepo := rsv_repo.NewRepository(db)
			userRepo := user_repo.NewRepository(db)
			rsvUC := rsv_uc.NewUsecase(rsvRepo, roomRepo)

			// ユーザーを作成
			_, err := userRepo.CreateUser(ctx, &user_repo.CreateUserArgs{
				ID:                "testuser1",
				EncryptedPassword: "testpass1",
			})
			require.NoError(t, err)

			_, err = userRepo.CreateUser(ctx, &user_repo.CreateUserArgs{
				ID:                "testuser2",
				EncryptedPassword: "testpass2",
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

			// testuser1 の予約を作成
			output, err := rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser1",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
				FromHour:   9,
				FromMinute: 30,
				ToHour:     12,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: nil,
			})
			require.NoError(t, err)
			require.Len(t, output.Reservations, 1)

			// 他のユーザー(testuser2)が予約を削除しようとする
			err = rsvUC.DeleteReservation(ctx, &input.DeleteReservation{
				Actor: actor.Actor{
					ID:   "testuser2",
					Role: actor.RoleUser,
				},
				ReservationID: output.Reservations[0].ID,
			})
			require.ErrorIs(t, err, errs.ErrNotYourReservation)
		})
	})
}
