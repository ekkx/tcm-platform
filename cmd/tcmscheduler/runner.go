package main

import (
	"context"
	"log/slog"
	"sync"
	"github.com/ekkx/tcm-platform/internal/app/assemble"
	"github.com/ekkx/tcm-platform/internal/app/gateway"
	"github.com/ekkx/tcm-platform/internal/config"
	"github.com/ekkx/tcm-platform/internal/domain/entity"
	"github.com/ekkx/tcm-platform/internal/domain/valueobject"
	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/platform/tcmutil"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
	"github.com/ekkx/tcm-platform/internal/platform/ymd"

	rsvad "github.com/ekkx/tcm-platform/internal/modules/reservation/adapter"
	rsvrepo "github.com/ekkx/tcm-platform/internal/modules/reservation/repository"
	userad "github.com/ekkx/tcm-platform/internal/modules/user/adapter"
	userrepo "github.com/ekkx/tcm-platform/internal/modules/user/repository"

	"github.com/ekkx/tcmrsv"
)

type ReservationResult struct {
	ReservationID  ulid.ULID
	OfficialSiteID string
	Success        bool
	Error          error
}

type Group struct {
	MasterUser   entity.User
	Reservations []*assemble.ReservationView
}

func Run(cfg *config.Config, clientFactory OfficialClientFactory, overrideDate *ymd.YMD) error {
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
	reservationCmd := rsvad.NewCommandAdapter(rsvRepo)
	userAsm := assemble.NewUserAssembler(userQuery)
	rsvAsm := assemble.NewReservationAssembler(reservationQuery, userAsm)

	// 対象日を決定する（overrideDate があればそれを使い、なければ二日後）
	targetDate := ymd.Today().AddDays(2)
	if overrideDate != nil {
		targetDate = *overrideDate
	}
	rsvIDs, err := reservationQuery.ListPendingReservationIDsByDate(ctx, targetDate)
	if err != nil {
		return err
	}

	if len(rsvIDs) == 0 {
		slog.Info("no pending reservations to process", slog.String("date", targetDate.String()))
		return nil
	}

	// 予約を扱いやすく組み立てる
	rsvListView, err := rsvAsm.BuildList(ctx, rsvIDs)
	if err != nil {
		return err
	}

	// 予約したユーザーのマスターユーザーごとに予約をグループ化
	groupMap := make(map[ulid.ULID]*Group)
	for _, rsvView := range rsvListView {
		masterUser := rsvView.UserView.MasterUser
		if masterUser == nil {
			masterUser = &rsvView.UserView.User
		}

		g, ok := groupMap[masterUser.ID]
		if !ok {
			g = &Group{MasterUser: *masterUser}
			groupMap[masterUser.ID] = g
		}
		g.Reservations = append(g.Reservations, rsvView)
	}

	slog.Info("processing reservations",
		slog.String("date", targetDate.String()),
		slog.Int("total_reservations", len(rsvListView)),
		slog.Int("user_groups", len(groupMap)),
	)

	// マスターユーザーごとに並行処理（セマフォで並行数を制限）
	maxJobs := cfg.Scheduler.MaxConcurrentJobs
	if maxJobs <= 0 {
		maxJobs = 1
	}
	sem := make(chan struct{}, maxJobs)

	var wg sync.WaitGroup
	for _, group := range groupMap {
		wg.Add(1)
		sem <- struct{}{} // セマフォ取得
		go func(g *Group) {
			defer wg.Done()
			defer func() { <-sem }() // セマフォ解放
			processGroup(ctx, g, clientFactory, reservationCmd)
		}(group)
	}
	wg.Wait()

	slog.Info("tcm-scheduler job completed")
	return nil
}

func processGroup(ctx context.Context, group *Group, clientFactory OfficialClientFactory, rsvCmd gateway.ReservationCommand) {
	master := group.MasterUser

	if master.OfficialSiteID == nil || master.OfficialSitePassword == nil {
		slog.Error("master user missing official site credentials",
			slog.String("user_id", master.ID.String()),
		)
		for _, rsv := range group.Reservations {
			markFailed(ctx, rsvCmd, rsv.Reservation.ID, "missing credentials")
		}
		return
	}

	// 公式サイトにログイン
	client := clientFactory()
	if err := client.Login(&tcmrsv.LoginParams{
		UserID:   *master.OfficialSiteID,
		Password: *master.OfficialSitePassword,
	}); err != nil {
		slog.Error("failed to login to official site",
			slog.String("user_id", master.ID.String()),
			slog.String("official_site_id", *master.OfficialSiteID),
			slog.String("error", err.Error()),
		)
		for _, rsv := range group.Reservations {
			markFailed(ctx, rsvCmd, rsv.Reservation.ID, err.Error())
		}
		return
	}

	slog.Info("logged in to official site",
		slog.String("official_site_id", *master.OfficialSiteID),
		slog.Int("reservations_count", len(group.Reservations)),
	)

	// 各予約を公式サイトに投入
	for _, rsvView := range group.Reservations {
		rsv := rsvView.Reservation

		date := rsv.Date
		tcmDate := tcmrsv.NewDate(date.Year, date.Month, date.Day)

		err := client.Reserve(&tcmrsv.ReserveParams{
			Campus:     tcmutil.ToTCMCampusType(rsv.CampusType),
			RoomID:     rsv.RoomID,
			Date:       tcmDate,
			FromHour:   rsv.FromHour,
			FromMinute: rsv.FromMinute,
			ToHour:     rsv.ToHour,
			ToMinute:   rsv.ToMinute,
		})

		if err != nil {
			slog.Error("reservation failed",
				slog.String("reservation_id", rsv.ID.String()),
				slog.String("error", err.Error()),
			)
			markFailed(ctx, rsvCmd, rsv.ID, err.Error())
			continue
		}

		slog.Info("reservation succeeded",
			slog.String("reservation_id", rsv.ID.String()),
			slog.String("room_id", rsv.RoomID),
			slog.String("date", rsv.Date.String()),
		)

		// 成功: ステータスを更新
		if err := rsvCmd.UpdateReservationStatus(ctx, &gateway.UpdateReservationStatusCommand{
			ID:     rsv.ID,
			Status: valueobject.ReservationStatusSuccess,
		}); err != nil {
			slog.Error("failed to update reservation status",
				slog.String("reservation_id", rsv.ID.String()),
				slog.String("error", err.Error()),
			)
		}
	}
}

func markFailed(ctx context.Context, rsvCmd gateway.ReservationCommand, reservationID ulid.ULID, reason string) {
	slog.Warn("marking reservation as failed",
		slog.String("reservation_id", reservationID.String()),
		slog.String("reason", reason),
	)
	if err := rsvCmd.UpdateReservationStatus(ctx, &gateway.UpdateReservationStatusCommand{
		ID:     reservationID,
		Status: valueobject.ReservationStatusFailed,
	}); err != nil {
		slog.Error("failed to update reservation status to failed",
			slog.String("reservation_id", reservationID.String()),
			slog.String("error", err.Error()),
		)
	}
}
