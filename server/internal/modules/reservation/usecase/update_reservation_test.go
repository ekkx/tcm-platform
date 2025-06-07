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

func TestUpdateReservation_正常系(t *testing.T) {
	t.Run("一般ユーザーが自分の予約を更新する", func(t *testing.T) {
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
			rsv, err := rsvRepo.GetReservationByID(ctx, output.Reservations[0].ID)
			require.NoError(t, err)
			require.NotNil(t, rsv)

			// 元の予約情報を保存
			originalRsv, err := rsvRepo.GetReservationByID(ctx, output.Reservations[0].ID)
			require.NoError(t, err)
			require.NotNil(t, originalRsv)

			// 予約を更新（全てのフィールドを変更）
			updatedBookerName := "Updated Booker"
			externalID := "EXT-123"
			updatedOutput, err := rsvUC.UpdateReservation(ctx, &input.UpdateReservation{
				Actor: actor.Actor{
					ID:   "testuser",
					Role: actor.RoleUser,
				},
				ID:         output.Reservations[0].ID,
				ExternalID: &externalID,                                     // 外部IDを追加
				CampusType: enum.CampusTypeNakameguro,                       // キャンパスを変更
				Date:       time.Date(2033, 10, 2, 4, 2, 3, 4, utils.JST()), // 日付を変更
				FromHour:   10,                                              // 開始時間を変更
				FromMinute: 0,                                               // 開始分を変更
				ToHour:     12,                                              // 終了時間を変更
				ToMinute:   30,                                              // 終了分を変更
				RoomID:     room.ID,                                         // 同じ部屋を使用
				BookerName: &updatedBookerName,                              // 予約者名を変更
			})
			require.NoError(t, err)
			require.NotNil(t, updatedOutput)

			// 予約が更新されたことを確認（全てのフィールドを検証）
			updatedRsv, err := rsvRepo.GetReservationByID(ctx, output.Reservations[0].ID)
			require.NoError(t, err)
			require.NotNil(t, updatedRsv)

			// 全フィールドの検証
			require.Equal(t, originalRsv.ID, updatedRsv.ID)                                                                              // IDは変わらない
			require.Equal(t, externalID, *updatedRsv.ExternalID)                                                                         // 外部IDが更新されている
			require.Equal(t, originalRsv.UserID, updatedRsv.UserID)                                                                      // ユーザーIDは変わらない
			require.Equal(t, enum.CampusTypeNakameguro, updatedRsv.CampusType)                                                           // キャンパスが更新されている
			require.Equal(t, room.ID, updatedRsv.RoomID)                                                                                 // 部屋IDは同じ
			require.Equal(t, time.Date(2033, 10, 2, 0, 0, 0, 0, utils.JST()).Format("2006-01-02"), updatedRsv.Date.Format("2006-01-02")) // 日付が更新されている
			require.Equal(t, 10, updatedRsv.FromHour)                                                                                    // 開始時間が更新されている
			require.Equal(t, 0, updatedRsv.FromMinute)                                                                                   // 開始分が更新されている
			require.Equal(t, 12, updatedRsv.ToHour)                                                                                      // 終了時間が更新されている
			require.Equal(t, 30, updatedRsv.ToMinute)                                                                                    // 終了分が更新されている
			require.Equal(t, updatedBookerName, *updatedRsv.BookerName)                                                                  // 予約者名が更新されている
			require.Equal(t, originalRsv.CreatedAt.Format(time.RFC3339), updatedRsv.CreatedAt.Format(time.RFC3339))                      // 作成日時は変わらない
		})
	})

	t.Run("システムユーザーが任意の予約を更新する", func(t *testing.T) {
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

			// 元の予約情報を保存
			originalRsv, err := rsvRepo.GetReservationByID(ctx, output.Reservations[0].ID)
			require.NoError(t, err)
			require.NotNil(t, originalRsv)

			// システムユーザーが予約を更新（全てのフィールドを変更）
			systemBookerName := "System Updated"
			systemExternalID := "SYS-456"
			updatedOutput, err := rsvUC.UpdateReservation(ctx, &input.UpdateReservation{
				Actor: actor.Actor{
					ID:   "testsystem",
					Role: actor.RoleSystem,
				},
				ID:         output.Reservations[0].ID,
				ExternalID: &systemExternalID,                               // 外部IDを追加
				CampusType: enum.CampusTypeNakameguro,                       // キャンパスを変更
				Date:       time.Date(2033, 11, 1, 4, 2, 3, 4, utils.JST()), // 月を変更
				FromHour:   13,                                              // 開始時間を変更
				FromMinute: 0,                                               // 開始分を変更
				ToHour:     14,                                              // 終了時間を変更
				ToMinute:   30,                                              // 終了分を変更
				RoomID:     room.ID,                                         // 同じ部屋を使用
				BookerName: &systemBookerName,                               // 予約者名を変更
			})
			require.NoError(t, err)
			require.NotNil(t, updatedOutput)

			// 予約が更新されたことを確認（全てのフィールドを検証）
			updatedRsv, err := rsvRepo.GetReservationByID(ctx, output.Reservations[0].ID)
			require.NoError(t, err)
			require.NotNil(t, updatedRsv)

			// 全フィールドの検証
			require.Equal(t, originalRsv.ID, updatedRsv.ID)                                                                              // IDは変わらない
			require.Equal(t, systemExternalID, *updatedRsv.ExternalID)                                                                   // 外部IDが更新されている
			require.Equal(t, originalRsv.UserID, updatedRsv.UserID)                                                                      // ユーザーIDは変わらない
			require.Equal(t, enum.CampusTypeNakameguro, updatedRsv.CampusType)                                                           // キャンパスが更新されている
			require.Equal(t, room.ID, updatedRsv.RoomID)                                                                                 // 部屋IDは同じ
			require.Equal(t, time.Date(2033, 11, 1, 0, 0, 0, 0, utils.JST()).Format("2006-01-02"), updatedRsv.Date.Format("2006-01-02")) // 日付が更新されている
			require.Equal(t, 13, updatedRsv.FromHour)                                                                                    // 開始時間が更新されている
			require.Equal(t, 0, updatedRsv.FromMinute)                                                                                   // 開始分が更新されている
			require.Equal(t, 14, updatedRsv.ToHour)                                                                                      // 終了時間が更新されている
			require.Equal(t, 30, updatedRsv.ToMinute)                                                                                    // 終了分が更新されている
			require.Equal(t, systemBookerName, *updatedRsv.BookerName)                                                                   // 予約者名が更新されている
			require.Equal(t, originalRsv.CreatedAt.Format(time.RFC3339), updatedRsv.CreatedAt.Format(time.RFC3339))                      // 作成日時は変わらない
		})
	})
}

func TestUpdateReservation_異常系(t *testing.T) {
	t.Run("存在しない予約の更新", func(t *testing.T) {
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

			nonExistentReservationID := 99999 // 存在しない予約ID

			// 予約が存在しないことを確認
			rsv, err := rsvRepo.GetReservationByID(ctx, nonExistentReservationID)
			require.NoError(t, err)
			require.Nil(t, rsv)

			// 存在しない予約を更新しようとする
			bookerName := "Test Booker"
			_, err = rsvUC.UpdateReservation(ctx, &input.UpdateReservation{
				Actor: actor.Actor{
					ID:   "testuser",
					Role: actor.RoleUser,
				},
				ID:         nonExistentReservationID,
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
				FromHour:   9,
				FromMinute: 30,
				ToHour:     12,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: &bookerName,
			})
			// 存在しない予約を更新しようとしてエラーになる
			require.ErrorIs(t, err, errs.ErrReservationNotFound)
		})
	})

	t.Run("他のユーザーの予約を更新しようとする", func(t *testing.T) {
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
			bookerName := "Test Booker 1"
			output, err := rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser1",
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

			// 他のユーザー(testuser2)が予約を更新しようとする
			updatedBookerName := "Updated by user2"
			_, err = rsvUC.UpdateReservation(ctx, &input.UpdateReservation{
				Actor: actor.Actor{
					ID:   "testuser2",
					Role: actor.RoleUser,
				},
				ID:         output.Reservations[0].ID,
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 2, 4, 2, 3, 4, utils.JST()),
				FromHour:   10,
				FromMinute: 0,
				ToHour:     12,
				ToMinute:   30,
				RoomID:     room.ID,
				BookerName: &updatedBookerName,
			})
			require.ErrorIs(t, err, errs.ErrNotYourReservation)

			// 予約が更新されていないことを確認
			rsv, err := rsvRepo.GetReservationByID(ctx, output.Reservations[0].ID)
			require.NoError(t, err)
			require.NotNil(t, rsv)
			require.Equal(t, bookerName, *rsv.BookerName) // 元の予約者名のままであることを確認
		})
	})

	t.Run("無効なパラメータでの更新", func(t *testing.T) {
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

			// 無効な時間範囲で更新を試みる（終了時間が開始時間より前）
			_, err = rsvUC.UpdateReservation(ctx, &input.UpdateReservation{
				Actor: actor.Actor{
					ID:   "testuser",
					Role: actor.RoleUser,
				},
				ID:         output.Reservations[0].ID,
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
				FromHour:   12, // 開始時間が終了時間より後
				FromMinute: 0,
				ToHour:     9,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: &bookerName,
			})
			require.ErrorIs(t, err, errs.ErrInvalidTimeRange)

			// 無効なキャンパスタイプで更新を試みる
			_, err = rsvUC.UpdateReservation(ctx, &input.UpdateReservation{
				Actor: actor.Actor{
					ID:   "testuser",
					Role: actor.RoleUser,
				},
				ID:         output.Reservations[0].ID,
				CampusType: enum.CampusType(999), // 無効なキャンパスタイプ
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
				FromHour:   9,
				FromMinute: 30,
				ToHour:     12,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: &bookerName,
			})
			require.ErrorIs(t, err, errs.ErrInvalidCampusType)
		})
	})

	t.Run("予約時間が他の予約と重複する場合", func(t *testing.T) {
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

			// 予約1を作成（9:00-11:00）
			bookerName1 := "Test Booker 1"
			output1, err := rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
				FromHour:   9,
				FromMinute: 0,
				ToHour:     11,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: &bookerName1,
			})
			require.NoError(t, err)
			require.Len(t, output1.Reservations, 1)

			// 予約2を作成（13:00-15:00）- 予約1と重複しない
			bookerName2 := "Test Booker 2"
			output2, err := rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
				FromHour:   13,
				FromMinute: 0,
				ToHour:     15,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: &bookerName2,
			})
			require.NoError(t, err)
			require.Len(t, output2.Reservations, 1)

			// 予約2を更新して予約1と重複させる（10:00-12:00）
			updatedBookerName := "Updated Booker"
			_, err = rsvUC.UpdateReservation(ctx, &input.UpdateReservation{
				Actor: actor.Actor{
					ID:   "testuser",
					Role: actor.RoleUser,
				},
				ID:         output2.Reservations[0].ID,
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
				FromHour:   10, // 予約1（9:00-11:00）と重複
				FromMinute: 0,
				ToHour:     12,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: &updatedBookerName,
			})
			// 予約が重複しているためエラーになる
			require.ErrorIs(t, err, errs.ErrReservationConflict)

			// 予約2が更新されていないことを確認
			rsv2, err := rsvRepo.GetReservationByID(ctx, output2.Reservations[0].ID)
			require.NoError(t, err)
			require.NotNil(t, rsv2)
			require.Equal(t, 13, rsv2.FromHour) // 元の時間のままであることを確認
			require.Equal(t, 0, rsv2.FromMinute)
			require.Equal(t, 15, rsv2.ToHour)
			require.Equal(t, 0, rsv2.ToMinute)
			require.Equal(t, bookerName2, *rsv2.BookerName) // 元の予約者名のままであることを確認
		})
	})
}
