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
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/ekkx/tcmrsv-web/server/pkg/utils"
	"github.com/ekkx/tcmrsv-web/server/tests/testhelper"
	"github.com/stretchr/testify/require"
)

func TestCreateReservation_正常系(t *testing.T) {
	// TODO: tcmrsv.GetRoomsFiltered のモックを作る

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

			// 予約の内容を確認
			require.Len(t, output.Reservations, 1)
			require.Equal(t, "testuser", output.Reservations[0].UserID)
			require.Equal(t, enum.CampusTypeIkebukuro, output.Reservations[0].CampusType)
			require.Equal(t, time.Date(2033, 10, 1, 0, 0, 0, 0, utils.JST()), output.Reservations[0].Date)
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
			_, err = rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
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

			// 予約を作成（開始時間と終了時間が同じ）
			_, err = rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
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

			// FromMinuteが不正な値（15分）
			_, err = rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
				FromHour:   9,
				FromMinute: 15,
				ToHour:     10,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: nil,
			})
			require.ErrorIs(t, err, errs.ErrInvalidArgument)

			// ToMinuteが不正な値（45分）
			_, err = rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
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

			// 無効なキャンパスタイプで予約を作成
			_, err = rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser",
				CampusType: 999, // 無効なキャンパスタイプ
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
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

			// 存在しないルームIDで予約を作成
			_, err = rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
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
			roomRepo := room_repo.NewRepository(tcmrsv.New())
			rsvRepo := rsv_repo.NewRepository(db)
			userRepo := user_repo.NewRepository(db)
			rsvUC := rsv_uc.NewUsecase(rsvRepo, roomRepo)

			// ユーザーを作成
			_, err := userRepo.CreateUser(ctx, &user_repo.CreateUserArgs{
				ID:                "testuser1",
				EncryptedPassword: "testpass",
			})
			require.NoError(t, err)

			_, err = userRepo.CreateUser(ctx, &user_repo.CreateUserArgs{
				ID:                "testuser2",
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

			// 最初の予約を作成
			_, err = rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser1",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
				FromHour:   9,
				FromMinute: 0,
				ToHour:     11,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: nil,
			})
			require.NoError(t, err)

			// 同じ時間帯で重複する予約を作成しようとする
			_, err = rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser2",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
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
			roomRepo := room_repo.NewRepository(tcmrsv.New())
			rsvRepo := rsv_repo.NewRepository(db)
			rsvUC := rsv_uc.NewUsecase(rsvRepo, roomRepo)

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

			// 存在しないユーザーIDで予約を作成
			_, err := rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "non-existent-user",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 1, 4, 2, 3, 4, utils.JST()),
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
