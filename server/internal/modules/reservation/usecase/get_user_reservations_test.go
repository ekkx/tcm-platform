package usecase_test

import (
	"testing"
	"time"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/reservation/dto/input"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/ekkx/tcmrsv-web/server/pkg/utils"
	"github.com/ekkx/tcmrsv-web/server/tests/testhelper"
	"github.com/stretchr/testify/require"
)

func TestGetUserReservations_正常系(t *testing.T) {
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

			// 予約を作成（全て未来の日付）
			now := time.Now().In(utils.JST())
			// 今日の0時を基準にする
			today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, utils.JST())
			tomorrow := today.Add(24 * time.Hour)
			dayAfterTomorrow := today.Add(48 * time.Hour)
			
			// 明日の予約
			rsv1 := testhelper.CreateTestReservation(ctx, t, deps.RsvUC, testhelper.TestUserID, room,
				tomorrow, 9, 0, 10, 0)
			
			// 明後日の予約
			rsv2 := testhelper.CreateTestReservation(ctx, t, deps.RsvUC, testhelper.TestUserID, room,
				dayAfterTomorrow, 14, 0, 15, 30)

			// 自分の予約を取得
			output, err := deps.RsvUC.GetUserReservations(ctx, &input.GetUserReservations{
				UserID: testhelper.TestUserID,
			})
			require.NoError(t, err)
			require.Len(t, output.Reservations, 2)
			
			// 予約が含まれていることを確認
			reservationIDs := make([]int, len(output.Reservations))
			for i, rsv := range output.Reservations {
				reservationIDs[i] = rsv.ID
			}
			require.Contains(t, reservationIDs, rsv1.ID)
			require.Contains(t, reservationIDs, rsv2.ID)
			
			// 日付順にソートされていることを確認
			for i := 1; i < len(output.Reservations); i++ {
				prev := output.Reservations[i-1]
				curr := output.Reservations[i]
				require.True(t, 
					prev.Date.Before(curr.Date) || 
					(prev.Date.Equal(curr.Date) && (prev.FromHour < curr.FromHour || 
						(prev.FromHour == curr.FromHour && prev.FromMinute <= curr.FromMinute))))
			}
		})
	})

	t.Run("特定の日付以降の予約を取得", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUser(ctx, t, deps.UserRepo)

			// ルームを取得
			room := testhelper.GetIkebukuroRoom(ctx, t, deps.RoomRepo)

			// 予約を作成（日付の開始時刻を明確に設定）
			now := time.Now().In(utils.JST())
			// 今日の0時を基準にする
			today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, utils.JST())
			tomorrow := today.Add(24 * time.Hour)
			dayAfterTomorrow := today.Add(48 * time.Hour)
			threeDaysLater := today.Add(72 * time.Hour)
			
			// 明日の予約
			testhelper.CreateTestReservation(ctx, t, deps.RsvUC, testhelper.TestUserID, room,
				tomorrow, 9, 0, 10, 0)
			
			// 明後日の予約
			rsv2 := testhelper.CreateTestReservation(ctx, t, deps.RsvUC, testhelper.TestUserID, room,
				dayAfterTomorrow, 10, 30, 12, 0)
			
			// 3日後の予約
			rsv3 := testhelper.CreateTestReservation(ctx, t, deps.RsvUC, testhelper.TestUserID, room,
				threeDaysLater, 14, 0, 15, 30)

			// 明後日以降の予約を取得
			output, err := deps.RsvUC.GetUserReservations(ctx, &input.GetUserReservations{
				UserID:   testhelper.TestUserID,
				FromDate: dayAfterTomorrow,
			})
			require.NoError(t, err)
			require.Len(t, output.Reservations, 2)
			
			// 明後日と3日後の予約が含まれていることを確認
			reservationIDs := make([]int, len(output.Reservations))
			for i, rsv := range output.Reservations {
				reservationIDs[i] = rsv.ID
			}
			require.Contains(t, reservationIDs, rsv2.ID)
			require.Contains(t, reservationIDs, rsv3.ID)
		})
	})

	t.Run("予約がない場合", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// ユーザーを作成
			testhelper.CreateDefaultTestUser(ctx, t, deps.UserRepo)

			// 予約を取得（予約なし）
			output, err := deps.RsvUC.GetUserReservations(ctx, &input.GetUserReservations{
				UserID: testhelper.TestUserID,
			})
			require.NoError(t, err)
			require.Empty(t, output.Reservations)
		})
	})
}

func TestGetUserReservations_異常系(t *testing.T) {
	t.Run("ユーザーが存在しない", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// 存在しないユーザーの予約を取得
			output, err := deps.RsvUC.GetUserReservations(ctx, &input.GetUserReservations{
				UserID: "non-existent-user",
			})
			// ユーザーが存在しなくても空の配列を返す
			require.NoError(t, err)
			require.Empty(t, output.Reservations)
		})
	})

	t.Run("過去の日付を指定", func(t *testing.T) {
		testhelper.RunWithTx(t, func(db database.Execer) {
			ctx := testhelper.GetContextWithConfig(t)

			// 依存関係のセットアップ
			deps := testhelper.SetupReservationTestDependencies(db)

			// 過去の日付で予約を取得しようとする
			_, err := deps.RsvUC.GetUserReservations(ctx, &input.GetUserReservations{
				UserID:   testhelper.TestUserID,
				FromDate: time.Now().Add(-24 * time.Hour), // 昨日
			})
			require.ErrorIs(t, err, errs.ErrDateMustBeTodayOrFuture)
		})
	})
}