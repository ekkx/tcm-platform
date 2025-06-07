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

func TestGetReservation_正常系(t *testing.T) {
	// TODO: tcmrsv.GetRoomsFiltered のモックを作る

	t.Run("自分の予約取得", func(t *testing.T) {
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

			// 予約内容の確認
			require.NotEmpty(t, rsv.ID)
			require.Nil(t, rsv.ExternalID)
			require.Equal(t, testhelper.TestUserID, rsv.UserID)
			require.Equal(t, room.CampusType, rsv.CampusType)
			require.Equal(t, room.ID, rsv.RoomID)
			require.Equal(t, testhelper.GetTestDate(), rsv.Date)
			require.Equal(t, 9, rsv.FromHour)
			require.Equal(t, 30, rsv.FromMinute)
			require.Equal(t, 12, rsv.ToHour)
			require.Equal(t, 0, rsv.ToMinute)
			require.Equal(t, bookerName, *rsv.BookerName)
			require.NotEmpty(t, rsv.CreatedAt)
		})
	})

	t.Run("システムがユーザーの予約を取得する", func(t *testing.T) {
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

			// システムが予約を取得
			output, err := deps.RsvUC.GetReservation(ctx, &input.GetReservation{
				Actor: actor.Actor{
					ID:   testhelper.TestSystemID,
					Role: actor.RoleSystem,
				},
				ReservationID: rsv.ID,
			})
			require.NoError(t, err)
			require.NotNil(t, output)
			require.Equal(t, rsv.ID, output.Reservation.ID)
		})
	})
}

func TestGetReservation_異常系(t *testing.T) {
	// TODO: tcmrsv.GetRoomsFiltered のモックを作る

	t.Run("予約が存在しない", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// 存在しない予約IDで取得を試みる
			_, err := deps.RsvUC.GetReservation(ctx, &input.GetReservation{
				Actor: actor.Actor{
					ID:   testhelper.TestUserID,
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
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUsers(ctx, t, deps.UserRepo)

			// ルームを取得
			room := testhelper.GetIkebukuroRoom(ctx, t, deps.RoomRepo)

			// 他人の予約を作成
			bookerName := "Test Booker"
			rsv := testhelper.CreateTestReservationWithParams(ctx, t, deps.RsvUC, testhelper.TestReservationParams{
				UserID:     testhelper.TestUserID2,
				Room:       room,
				Date:       testhelper.GetTestDateTime(),
				FromHour:   9,
				FromMinute: 30,
				ToHour:     12,
				ToMinute:   0,
				BookerName: &bookerName,
			})

			// 他人の予約を取得しようとする
			_, err := deps.RsvUC.GetReservation(ctx, &input.GetReservation{
				Actor: actor.Actor{
					ID:   testhelper.TestUserID,
					Role: actor.RoleUser,
				},
				ReservationID: rsv.ID,
			})
			require.Error(t, err)
			require.ErrorIs(t, err, errs.ErrNotYourReservation)
		})
	})
}