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
	"github.com/ekkx/tcmrsv-web/server/internal/shared/errs"
)

func TestGetUserReservations_正常系(t *testing.T) {
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

			output2, err := rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 2, 4, 2, 3, 4, utils.JST()),
				FromHour:   13,
				FromMinute: 00,
				ToHour:     14,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: &bookerName,
			})
			require.NoError(t, err)

			rsvsOutput, err := rsvUC.GetUserReservations(ctx, &input.GetUserReservations{
				UserID: "testuser",
			})
			require.NoError(t, err)

			for i, rsv := range rsvsOutput.Reservations {
				require.NotEmpty(t, rsv.ID)
				require.Nil(t, rsv.ExternalID)
				require.Equal(t, "testuser", rsv.UserID)
				require.Equal(t, enum.CampusTypeIkebukuro, rsv.CampusType)
				require.Equal(t, room.ID, rsv.RoomID)
				if i == 0 {
					require.Equal(t, time.Date(2033, 10, 1, 0, 0, 0, 0, utils.JST()), rsv.Date)
					require.Equal(t, output.Reservations[0].FromHour, rsv.FromHour)
					require.Equal(t, output.Reservations[0].FromMinute, rsv.FromMinute)
					require.Equal(t, output.Reservations[0].ToHour, rsv.ToHour)
					require.Equal(t, output.Reservations[0].ToMinute, rsv.ToMinute)
				} else {
					require.Equal(t, time.Date(2033, 10, 2, 0, 0, 0, 0, utils.JST()), rsv.Date)
					require.Equal(t, output2.Reservations[0].FromHour, rsv.FromHour)
					require.Equal(t, output2.Reservations[0].FromMinute, rsv.FromMinute)
					require.Equal(t, output2.Reservations[0].ToHour, rsv.ToHour)
					require.Equal(t, output2.Reservations[0].ToMinute, rsv.ToMinute)
				}
				require.Equal(t, bookerName, *rsv.BookerName)
				require.NotEmpty(t, rsv.CreatedAt)
			}
		})
	})

	t.Run("日付を指定して自分の予約取得", func(t *testing.T) {
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
			_, err = rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 9, 30, 4, 2, 3, 4, utils.JST()),
				FromHour:   9,
				FromMinute: 30,
				ToHour:     12,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: &bookerName,
			})
			require.NoError(t, err)

			output2, err := rsvUC.CreateReservation(ctx, &input.CreateReservation{
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

			output3, err := rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(2033, 10, 2, 4, 2, 3, 4, utils.JST()),
				FromHour:   13,
				FromMinute: 0,
				ToHour:     14,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: &bookerName,
			})
			require.NoError(t, err)

			rsvsOutput, err := rsvUC.GetUserReservations(ctx, &input.GetUserReservations{
				UserID:   "testuser",
				FromDate: time.Date(2033, 10, 1, 0, 0, 0, 0, utils.JST()),
			})
			require.NoError(t, err)
			require.Len(t, rsvsOutput.Reservations, 2)

			for i, rsv := range rsvsOutput.Reservations {
				require.Equal(t, "testuser", rsv.UserID)
				if i == 0 {
					require.Equal(t, output2.Reservations[0].ID, rsv.ID)
					require.Equal(t, time.Date(2033, 10, 1, 0, 0, 0, 0, utils.JST()), rsv.Date)
					require.Equal(t, output2.Reservations[0].FromHour, rsv.FromHour)
					require.Equal(t, output2.Reservations[0].FromMinute, rsv.FromMinute)
					require.Equal(t, output2.Reservations[0].ToHour, rsv.ToHour)
					require.Equal(t, output2.Reservations[0].ToMinute, rsv.ToMinute)
				} else {
					require.Equal(t, output3.Reservations[0].ID, rsv.ID)
					require.Equal(t, time.Date(2033, 10, 2, 0, 0, 0, 0, utils.JST()), rsv.Date)
					require.Equal(t, output3.Reservations[0].FromHour, rsv.FromHour)
					require.Equal(t, output3.Reservations[0].FromMinute, rsv.FromMinute)
					require.Equal(t, output3.Reservations[0].ToHour, rsv.ToHour)
					require.Equal(t, output3.Reservations[0].ToMinute, rsv.ToMinute)
				}
			}

			rsvsOutput2, err := rsvUC.GetUserReservations(ctx, &input.GetUserReservations{
				UserID:   "testuser",
				FromDate: time.Date(2033, 10, 2, 0, 0, 0, 0, utils.JST()),
			})
			require.NoError(t, err)
			require.Len(t, rsvsOutput2.Reservations, 1)
			require.Equal(t, "testuser", rsvsOutput2.Reservations[0].UserID)
			require.Equal(t, output3.Reservations[0].ID, rsvsOutput2.Reservations[0].ID)
			require.Equal(t, time.Date(2033, 10, 2, 0, 0, 0, 0, utils.JST()), rsvsOutput2.Reservations[0].Date)
			require.Equal(t, output3.Reservations[0].FromHour, rsvsOutput2.Reservations[0].FromHour)
			require.Equal(t, output3.Reservations[0].FromMinute, rsvsOutput2.Reservations[0].FromMinute)
			require.Equal(t, output3.Reservations[0].ToHour, rsvsOutput2.Reservations[0].ToHour)
			require.Equal(t, output3.Reservations[0].ToMinute, rsvsOutput2.Reservations[0].ToMinute)
		})
	})

	t.Run("予約が存在しない", func(t *testing.T) {
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

			rsvsOutput, err := rsvUC.GetUserReservations(ctx, &input.GetUserReservations{
				UserID: "testuser",
			})
			require.NoError(t, err)
			require.Len(t, rsvsOutput.Reservations, 0)
		})
	})
}

func TestGetUserReservations_異常系(t *testing.T) {
	// TODO: tcmrsv.GetRoomsFiltered のモックを作る

	t.Run("過去の日付を指定して取得", func(t *testing.T) {
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

			today := time.Now().In(utils.JST())

			// 予約を作成
			bookerName := "Test Booker"
			_, err = rsvUC.CreateReservation(ctx, &input.CreateReservation{
				UserID:     "testuser",
				CampusType: enum.CampusTypeIkebukuro,
				Date:       time.Date(today.Year(), today.Month(), today.Day(), 4, 2, 3, 4, utils.JST()).AddDate(0, 0, -1), // 昨日の日付
				FromHour:   9,
				FromMinute: 30,
				ToHour:     12,
				ToMinute:   0,
				RoomID:     room.ID,
				BookerName: &bookerName,
			})
			require.NoError(t, err)

			// 過去の日付を指定して予約を取得
			_, err = rsvUC.GetUserReservations(ctx, &input.GetUserReservations{
				UserID:   "testuser",
				FromDate: time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, utils.JST()).AddDate(0, 0, -1), // 昨日の日付
			})
			require.ErrorIs(t, err, errs.ErrInvalidArgument)
		})
	})
}
