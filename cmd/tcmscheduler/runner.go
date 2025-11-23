package main

import (
	"context"
	"log/slog"
	"sync"

	"github.com/ekkx/tcm-platform/internal/app/assemble"
	"github.com/ekkx/tcm-platform/internal/config"
	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
	"github.com/ekkx/tcm-platform/internal/platform/ymd"

	rsvad "github.com/ekkx/tcm-platform/internal/modules/reservation/adapter"
	rsvrepo "github.com/ekkx/tcm-platform/internal/modules/reservation/repository"
	userad "github.com/ekkx/tcm-platform/internal/modules/user/adapter"
	userrepo "github.com/ekkx/tcm-platform/internal/modules/user/repository"
)

type ReservationResult struct {
	RetryCount     int
	OfficialSiteID string
	Reservation    assemble.ReservationView
}

type Group struct {
	MasterUser   entity.User
	Reservations []*assemble.ReservationView
	Results      []ReservationResult
}

func Run(cfg *config.Config) error {
	slog.Info("tcm-scheduler job started")

	ctx := context.Background()
	pool, err := cfg.Database.Open(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()

	// 依存関係の初期化
	querier := sqlc.New(pool)
	userRepo := userrepo.New(querier)
	rsvRepo := rsvrepo.New(querier)
	userQuery := userad.NewQueryAdapter(userRepo)
	reservationQuery := rsvad.NewQueryAdapter(rsvRepo)
	userAsm := assemble.NewUserAssembler(userQuery)
	rsvAsm := assemble.NewReservationAssembler(reservationQuery, userAsm)

	// 二日後の予約を取得する
	rsvIDs, err := reservationQuery.ListReservationIDsByDate(ctx, ymd.Today().AddDays(2))
	if err != nil {
		return err
	}

	// 予約を扱いやすく組み立てる
	rsvListView, err := rsvAsm.BuildList(ctx, rsvIDs)
	if err != nil {
		return err
	}

	// 予約したユーザーのマスターユーザーごとに予約をグループ化（map masterUserID  reservations）
	rsvGroupsByMasterUser := make(map[ulid.ULID][]Group)
	for _, rsvView := range rsvListView {
		// マスターユーザーを抽出する
		masterUser := rsvView.UserView.MasterUser
		if masterUser == nil {
			masterUser = &rsvView.UserView.User
		}
		// マスターユーザーごとに予約をグループ化する
		slog.Debug("adding reservation to user group", slog.String("user_id", masterUser.ID.String()), slog.String("reservation_id", rsvView.Reservation.ID.String()))
		rsvGroupsByMasterUser[masterUser.ID] = append(rsvGroupsByMasterUser[masterUser.ID], Group{
			MasterUser:   *masterUser,
			Reservations: []*assemble.ReservationView{rsvView},
		})
	}

	// 予約がない場合は終了
	if len(rsvGroupsByMasterUser) == 0 {
		slog.Info("no reservations to process")
		return nil
	}

	slog.Debug("grouped reservations", slog.Int("user_groups_count", len(rsvGroupsByMasterUser)))

	// await reservation processing results and retry if failed
	var wg sync.WaitGroup
	for _, groups := range rsvGroupsByMasterUser {
		slog.Info("processsing user group", slog.String("master_user_id", groups[0].MasterUser.ID.String()), slog.Int("groups_count", len(groups)))

		for _, group := range groups {
			slog.Info("processing reservation group", slog.String("master_user_id", group.MasterUser.ID.String()), slog.Int("reservations_count", len(group.Reservations)))
			wg.Add(1)
			go func(group Group) {
				defer wg.Done()
				processReservations(group.MasterUser, group.Reservations)
			}(group)
		}
	}
	wg.Wait()

	return nil
}

func processReservations(masterUser entity.User, reservationList []*assemble.ReservationView) {
	slog.Debug("official site login", slog.String("official_site_id", *masterUser.OfficialSiteID), slog.String("official_site_password", *masterUser.OfficialSitePassword))
	for _, rsv := range reservationList {
		slog.Debug("processing reservation", slog.String("reservation_id", rsv.Reservation.ID.String()), slog.String("date", rsv.Reservation.Date.String()), slog.String("room_id", rsv.Reservation.RoomID))
	}
}
