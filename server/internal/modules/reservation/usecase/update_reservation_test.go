package usecase_test

import (
	"testing"

	"github.com/ekkx/tcmrsv-web/server/internal/domain/entity"
	"github.com/ekkx/tcmrsv-web/server/internal/domain/enum"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	room_repo "github.com/ekkx/tcmrsv-web/server/internal/modules/room/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/actor"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/ekkx/tcmrsv-web/server/tests/testhelper"
	"github.com/stretchr/testify/require"
)

func TestUpdateReservation_正常系(t *testing.T) {
	t.Run("一般ユーザーが自分の予約を更新する", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUser(ctx, t, deps.UserRepo)

			// ルームを取得
			room := testhelper.GetIkebukuroRoom(ctx, t, deps.RoomRepo)

			// 予約を作成
			rsv := testhelper.CreateDefaultTestReservation(ctx, t, deps.RsvUC, testhelper.TestUserID, room)

			// 予約の時間を変更
			output, err := deps.RsvUC.UpdateReservation(ctx, &input.UpdateReservation{
				Actor: actor.Actor{
					ID:   testhelper.TestUserID,
					Role: actor.RoleUser,
				},
				ID:         rsv.ID,
				CampusType: enum.CampusTypeIkebukuro,
				Date:       testhelper.GetTestDateTime(),
				FromHour:   14, // 変更: 9:30 -> 14:00
				FromMinute: 0,
				ToHour:     16, // 変更: 12:00 -> 16:00
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: nil,
			})
			require.NoError(t, err)

			// 更新内容の確認
			require.Equal(t, rsv.ID, output.Reservation.ID)
			require.Equal(t, testhelper.TestUserID, output.Reservation.UserID)
			require.Equal(t, enum.CampusTypeIkebukuro, output.Reservation.CampusType)
			require.Equal(t, testhelper.GetTestDate(), output.Reservation.Date)
			require.Equal(t, 14, output.Reservation.FromHour)
			require.Equal(t, 0, output.Reservation.FromMinute)
			require.Equal(t, 16, output.Reservation.ToHour)
			require.Equal(t, 0, output.Reservation.ToMinute)
			require.Equal(t, room.ID, output.Reservation.RoomID)
			require.Nil(t, output.Reservation.BookerName)
		})
	})

	t.Run("システムが予約を更新する", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUser(ctx, t, deps.UserRepo)

			// ルームを取得
			room := testhelper.GetIkebukuroRoom(ctx, t, deps.RoomRepo)

			// 予約を作成
			rsv := testhelper.CreateDefaultTestReservation(ctx, t, deps.RsvUC, testhelper.TestUserID, room)

			// システムが予約を更新
			updatedBookerName := "Updated by System"
			output, err := deps.RsvUC.UpdateReservation(ctx, &input.UpdateReservation{
				Actor: actor.Actor{
					ID:   testhelper.TestSystemID,
					Role: actor.RoleSystem,
				},
				ID:         rsv.ID,
				CampusType: enum.CampusTypeIkebukuro,
				Date:       testhelper.GetTestDateTime(),
				FromHour:   9,
				FromMinute: 30,
				ToHour:     12,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: &updatedBookerName,
			})
			require.NoError(t, err)

			// 更新内容の確認
			require.Equal(t, updatedBookerName, *output.Reservation.BookerName)
		})
	})

	t.Run("違う部屋に変更する", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUser(ctx, t, deps.UserRepo)

			// ルームを取得
			rooms := deps.RoomRepo.SearchRooms(ctx, &room_repo.SearchRoomsArgs{})
			require.GreaterOrEqual(t, len(rooms), 2)

			// 2つの異なる部屋を取得
			var room1, room2 entity.Room
			foundFirst := false
			for _, r := range rooms {
				if r.CampusType == enum.CampusTypeIkebukuro {
					if !foundFirst {
						room1 = r
						foundFirst = true
					} else if r.ID != room1.ID {
						room2 = r
						break
					}
				}
			}
			require.NotEmpty(t, room1.ID)
			require.NotEmpty(t, room2.ID)

			// 最初の部屋で予約を作成
			rsv := testhelper.CreateDefaultTestReservation(ctx, t, deps.RsvUC, testhelper.TestUserID, room1)

			// 別の部屋に変更
			output, err := deps.RsvUC.UpdateReservation(ctx, &input.UpdateReservation{
				Actor: actor.Actor{
					ID:   testhelper.TestUserID,
					Role: actor.RoleUser,
				},
				ID:         rsv.ID,
				CampusType: enum.CampusTypeIkebukuro,
				Date:       testhelper.GetTestDateTime(),
				FromHour:   9,
				FromMinute: 30,
				ToHour:     12,
				ToMinute:   0,
				RoomID:     room2.ID,
				BookerName: nil,
			})
			require.NoError(t, err)

			// 部屋が変更されたことを確認
			require.Equal(t, room2.ID, output.Reservation.RoomID)
		})
	})
}

func TestUpdateReservation_異常系(t *testing.T) {
	t.Run("存在しない予約を更新する", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUser(ctx, t, deps.UserRepo)

			// ルームを取得
			room := testhelper.GetIkebukuroRoom(ctx, t, deps.RoomRepo)

			// 存在しない予約を更新
			_, err := deps.RsvUC.UpdateReservation(ctx, &input.UpdateReservation{
				Actor: actor.Actor{
					ID:   testhelper.TestUserID,
					Role: actor.RoleUser,
				},
				ID:         99999, // 存在しないID
				CampusType: enum.CampusTypeIkebukuro,
				Date:       testhelper.GetTestDateTime(),
				FromHour:   9,
				FromMinute: 30,
				ToHour:     12,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: nil,
			})
			require.ErrorIs(t, err, errs.ErrReservationNotFound)
		})
	})

	t.Run("他のユーザーの予約を更新する", func(t *testing.T) {
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

			// testuser2 が予約を更新しようとする
			_, err := deps.RsvUC.UpdateReservation(ctx, &input.UpdateReservation{
				Actor: actor.Actor{
					ID:   testhelper.TestUserID2,
					Role: actor.RoleUser,
				},
				ID:         rsv.ID,
				CampusType: enum.CampusTypeIkebukuro,
				Date:       testhelper.GetTestDateTime(),
				FromHour:   14,
				FromMinute: 0,
				ToHour:     16,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: nil,
			})
			require.ErrorIs(t, err, errs.ErrNotYourReservation)
		})
	})

	t.Run("他の予約と重複する時間に更新する", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUsers(ctx, t, deps.UserRepo)

			// ルームを取得
			room := testhelper.GetIkebukuroRoom(ctx, t, deps.RoomRepo)

			// testuser の予約を作成（9:30-12:00）
			rsv1 := testhelper.CreateDefaultTestReservation(ctx, t, deps.RsvUC, testhelper.TestUserID, room)

			// testuser2 の予約を作成（14:00-16:00）
			testhelper.CreateTestReservation(ctx, t, deps.RsvUC, testhelper.TestUserID2, room,
				testhelper.GetTestDateTime(), 14, 0, 16, 0)

			// testuser が自分の予約を他の予約と重複する時間に更新しようとする
			_, err := deps.RsvUC.UpdateReservation(ctx, &input.UpdateReservation{
				Actor: actor.Actor{
					ID:   testhelper.TestUserID,
					Role: actor.RoleUser,
				},
				ID:         rsv1.ID,
				CampusType: enum.CampusTypeIkebukuro,
				Date:       testhelper.GetTestDateTime(),
				FromHour:   13, // 13:00-15:00 に変更しようとする（14:00-16:00と重複）
				FromMinute: 0,
				ToHour:     15,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: nil,
			})
			require.ErrorIs(t, err, errs.ErrReservationConflict)
		})
	})

	t.Run("不正な時間範囲に更新する", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUser(ctx, t, deps.UserRepo)

			// ルームを取得
			room := testhelper.GetIkebukuroRoom(ctx, t, deps.RoomRepo)

			// 予約を作成
			rsv := testhelper.CreateDefaultTestReservation(ctx, t, deps.RsvUC, testhelper.TestUserID, room)

			// 開始時間が終了時間より後に更新しようとする
			_, err := deps.RsvUC.UpdateReservation(ctx, &input.UpdateReservation{
				Actor: actor.Actor{
					ID:   testhelper.TestUserID,
					Role: actor.RoleUser,
				},
				ID:         rsv.ID,
				CampusType: enum.CampusTypeIkebukuro,
				Date:       testhelper.GetTestDateTime(),
				FromHour:   16,
				FromMinute: 0,
				ToHour:     14,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: nil,
			})
			require.ErrorIs(t, err, errs.ErrInvalidTimeRange)
		})
	})
}
