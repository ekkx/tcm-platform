package usecase_test

import (
	"testing"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/ekkx/tcmrsv-web/server/tests/testhelper"
	"github.com/stretchr/testify/require"
)

func TestDeleteReservation_正常系(t *testing.T) {
	// TODO: tcmrsv.GetRoomsFiltered のモックを作る

	t.Run("予約削除", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUser(ctx, t, deps.UserRepo)

			// ルームを取得
			room := testhelper.GetIkebukuroRoom(ctx, t, deps.RoomRepo)

			// 予約を作成
			bookerName := "Test Booker"
			rsv := testhelper.CreateTestReservationWithParams(ctx, t, deps.RsvUC, testhelper.TestReservationParams{
				UserID:     testhelper.TestUserID,
				Room:       room,
				Date:       testhelper.GetTestDateTime(),
				FromHour:   9,
				FromMinute: 30,
				ToHour:     12,
				ToMinute:   0,
				BookerName: &bookerName,
			})

			// 予約が作成されたことを確認
			_, err := deps.RsvRepo.GetReservationByID(ctx, rsv.ID)
			require.NoError(t, err)

			// 予約を削除
			err = deps.RsvUC.DeleteReservation(ctx, &input.DeleteReservation{
				Actor: actor.Actor{
					ID:   testhelper.TestUserID,
					Role: actor.RoleUser,
				},
				ReservationID: rsv.ID,
			})
			require.NoError(t, err)

			// 予約が削除されたことを確認
			deletedRsv, err := deps.RsvRepo.GetReservationByID(ctx, rsv.ID)
			require.NoError(t, err)
			require.Nil(t, deletedRsv)
		})
	})

	t.Run("システムユーザーによる予約削除", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUser(ctx, t, deps.UserRepo)

			// ルームを取得
			room := testhelper.GetIkebukuroRoom(ctx, t, deps.RoomRepo)

			// testuser の予約を作成
			rsv := testhelper.CreateDefaultTestReservation(ctx, t, deps.RsvUC, testhelper.TestUserID, room)

			// システムが予約を削除
			err := deps.RsvUC.DeleteReservation(ctx, &input.DeleteReservation{
				Actor: actor.Actor{
					ID:   testhelper.TestSystemID,
					Role: actor.RoleSystem,
				},
				ReservationID: rsv.ID,
			})
			require.NoError(t, err)

			// 予約が削除されたことを確認
			deletedRsv, err := deps.RsvRepo.GetReservationByID(ctx, rsv.ID)
			require.NoError(t, err)
			require.Nil(t, deletedRsv)
		})
	})
}

func TestDeleteReservation_異常系(t *testing.T) {
	// TODO: tcmrsv.GetRoomsFiltered のモックを作る

	t.Run("存在しない予約の削除", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUser(ctx, t, deps.UserRepo)

			nonExistentReservationID := 99999 // 存在しない予約ID

			// 予約が存在しないことを確認
			rsv, err := deps.RsvRepo.GetReservationByID(ctx, nonExistentReservationID)
			require.NoError(t, err)
			require.Nil(t, rsv)

			// 予約を削除 (user として)
			err = deps.RsvUC.DeleteReservation(ctx, &input.DeleteReservation{
				Actor: actor.Actor{
					ID:   testhelper.TestUserID,
					Role: actor.RoleUser,
				},
				ReservationID: nonExistentReservationID,
			})
			require.ErrorIs(t, err, errs.ErrReservationNotFound)

			// 予約を削除 (system として)
			err = deps.RsvUC.DeleteReservation(ctx, &input.DeleteReservation{
				Actor: actor.Actor{
					ID:   testhelper.TestSystemID,
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
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUsers(ctx, t, deps.UserRepo)

			// ルームを取得
			room := testhelper.GetIkebukuroRoom(ctx, t, deps.RoomRepo)

			// testuser の予約を作成
			rsv := testhelper.CreateDefaultTestReservation(ctx, t, deps.RsvUC, testhelper.TestUserID, room)

			// 他のユーザー(testuser2)が予約を削除しようとする
			err := deps.RsvUC.DeleteReservation(ctx, &input.DeleteReservation{
				Actor: actor.Actor{
					ID:   testhelper.TestUserID2,
					Role: actor.RoleUser,
				},
				ReservationID: rsv.ID,
			})
			require.ErrorIs(t, err, errs.ErrNotYourReservation)
		})
	})
}