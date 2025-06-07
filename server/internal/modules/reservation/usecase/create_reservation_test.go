package usecase_test

import (
	"testing"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/ekkx/tcmrsv-web/server/tests/testhelper"
	"github.com/stretchr/testify/require"
)

func TestCreateReservation_正常系(t *testing.T) {
	// TODO: tcmrsv.GetRoomsFiltered のモックを作る

	t.Run("練習室を指定して予約作成", func(t *testing.T) {
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
			output, err := deps.RsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     testhelper.TestUserID,
				CampusType: enum.CampusTypeIkebukuro,
				Date:       testhelper.GetTestDateTime(),
				FromHour:   9,
				FromMinute: 30,
				ToHour:     12,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: &bookerName,
			})
			require.NoError(t, err)

			// 予約の内容を確認
			require.Len(t, output.Reservations, 1)
			require.Equal(t, testhelper.TestUserID, output.Reservations[0].UserID)
			require.Equal(t, enum.CampusTypeIkebukuro, output.Reservations[0].CampusType)
			require.Equal(t, testhelper.GetTestDate(), output.Reservations[0].Date)
			require.Equal(t, 9, output.Reservations[0].FromHour)
			require.Equal(t, 30, output.Reservations[0].FromMinute)
			require.Equal(t, 12, output.Reservations[0].ToHour)
			require.Equal(t, 0, output.Reservations[0].ToMinute)
			require.Equal(t, room.ID, output.Reservations[0].RoomID)
			require.Equal(t, bookerName, *output.Reservations[0].BookerName)
		})
	})
}

func TestCreateReservation_異常系(t *testing.T) {
	// TODO: tcmrsv.GetRoomsFiltered のモックを作る

	t.Run("予約の開始時間が終了時刻よりも後", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUser(ctx, t, deps.UserRepo)

			// ルームを取得
			room := testhelper.GetIkebukuroRoom(ctx, t, deps.RoomRepo)

			// 予約を作成
			_, err := deps.RsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     testhelper.TestUserID,
				CampusType: enum.CampusTypeIkebukuro,
				Date:       testhelper.GetTestDateTime(),
				FromHour:   14,
				FromMinute: 0,
				ToHour:     12,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: nil,
			})
			require.ErrorIs(t, err, errs.ErrInvalidTimeRange)
		})
	})

	t.Run("予約の開始時間が終了時刻と同じ", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUser(ctx, t, deps.UserRepo)

			// ルームを取得
			room := testhelper.GetIkebukuroRoom(ctx, t, deps.RoomRepo)

			// 予約を作成（開始時間と終了時間が同じ）
			_, err := deps.RsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     testhelper.TestUserID,
				CampusType: enum.CampusTypeIkebukuro,
				Date:       testhelper.GetTestDateTime(),
				FromHour:   14,
				FromMinute: 30,
				ToHour:     14,
				ToMinute:   30,
				RoomID:     room.ID,
				BookerName: nil,
			})
			require.ErrorIs(t, err, errs.ErrInvalidTimeRange)
		})
	})

	t.Run("予約の分単位が不正（0か30以外）", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUser(ctx, t, deps.UserRepo)

			// ルームを取得
			room := testhelper.GetIkebukuroRoom(ctx, t, deps.RoomRepo)

			// FromMinuteが不正な値（15分）
			_, err := deps.RsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     testhelper.TestUserID,
				CampusType: enum.CampusTypeIkebukuro,
				Date:       testhelper.GetTestDateTime(),
				FromHour:   9,
				FromMinute: 15,
				ToHour:     10,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: nil,
			})
			require.ErrorIs(t, err, errs.ErrInvalidArgument)

			// ToMinuteが不正な値（45分）
			_, err = deps.RsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     testhelper.TestUserID,
				CampusType: enum.CampusTypeIkebukuro,
				Date:       testhelper.GetTestDateTime(),
				FromHour:   9,
				FromMinute: 0,
				ToHour:     10,
				ToMinute:   45,
				RoomID:     room.ID,
				BookerName: nil,
			})
			require.ErrorIs(t, err, errs.ErrInvalidArgument)
		})
	})

	t.Run("キャンパスが無効", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUser(ctx, t, deps.UserRepo)

			// 無効なキャンパスタイプで予約を作成
			_, err := deps.RsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     testhelper.TestUserID,
				CampusType: 999, // 無効なキャンパスタイプ
				Date:       testhelper.GetTestDateTime(),
				FromHour:   9,
				FromMinute: 0,
				ToHour:     10,
				ToMinute:   0,
				RoomID:     "room-123",
				BookerName: nil,
			})
			require.ErrorIs(t, err, errs.ErrInvalidCampusType)
		})
	})

	t.Run("練習室が存在しない", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUser(ctx, t, deps.UserRepo)

			// 存在しないルームIDで予約を作成
			_, err := deps.RsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     testhelper.TestUserID,
				CampusType: enum.CampusTypeIkebukuro,
				Date:       testhelper.GetTestDateTime(),
				FromHour:   9,
				FromMinute: 0,
				ToHour:     10,
				ToMinute:   0,
				RoomID:     "non-existent-room-id",
				BookerName: nil,
			})
			require.ErrorIs(t, err, errs.ErrRoomNotFound)
		})
	})

	t.Run("予約の重複", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUsers(ctx, t, deps.UserRepo)

			// ルームを取得
			room := testhelper.GetIkebukuroRoom(ctx, t, deps.RoomRepo)

			// 最初の予約を作成
			_, err := deps.RsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     testhelper.TestUserID,
				CampusType: enum.CampusTypeIkebukuro,
				Date:       testhelper.GetTestDateTime(),
				FromHour:   9,
				FromMinute: 0,
				ToHour:     11,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: nil,
			})
			require.NoError(t, err)

			// 同じ時間帯で重複する予約を作成しようとする
			_, err = deps.RsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     testhelper.TestUserID2,
				CampusType: enum.CampusTypeIkebukuro,
				Date:       testhelper.GetTestDateTime(),
				FromHour:   10,
				FromMinute: 0,
				ToHour:     12,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: nil,
			})
			require.ErrorIs(t, err, errs.ErrReservationConflict)
		})
	})

	t.Run("予約のユーザーが存在しない", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// ルームを取得
			room := testhelper.GetIkebukuroRoom(ctx, t, deps.RoomRepo)

			// 存在しないユーザーIDで予約を作成
			_, err := deps.RsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "non-existent-user",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       testhelper.GetTestDateTime(),
				FromHour:   9,
				FromMinute: 0,
				ToHour:     10,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: nil,
			})
			// ユーザーが存在しない場合、データベースの外部キー制約でエラーになる
			require.Error(t, err)
		})
	})
}
