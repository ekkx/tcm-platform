package main

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
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

// ReservationItemResult は1件の予約処理結果
type ReservationItemResult struct {
	RoomID     string
	Date       string
	FromHour   int
	FromMinute int
	ToHour     int
	ToMinute   int
	Success    bool
	Err        error
}

// GroupResult はマスターユーザー単位の処理結果
type GroupResult struct {
	DisplayID    string // 公式サイトID（未設定の場合は内部ユーザーID）
	MissingCreds bool
	LoginErr     error
	Items        []ReservationItemResult
}

func (g GroupResult) succeeded() int {
	n := 0
	for _, item := range g.Items {
		if item.Success {
			n++
		}
	}
	return n
}

func (g GroupResult) failed() int {
	if g.MissingCreds || g.LoginErr != nil {
		return len(g.Items)
	}
	n := 0
	for _, item := range g.Items {
		if !item.Success {
			n++
		}
	}
	return n
}

type Group struct {
	MasterUser   entity.User
	Reservations []*assemble.ReservationView
}

func Run(cfg *config.Config, clientFactory OfficialClientFactory, notifier Notifier, overrideDate *ymd.YMD) error {
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
		notifier.Notify(NotifyMessage{
			Level: NotifyLevelDebug,
			Title: "処理対象の予約なし",
			Body:  fmt.Sprintf("日付 %s の未処理予約はありませんでした。", targetDate.String()),
		})
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

	var (
		wg      sync.WaitGroup
		mu      sync.Mutex
		results []GroupResult
	)

	for _, group := range groupMap {
		wg.Add(1)
		sem <- struct{}{}
		go func(g *Group) {
			defer wg.Done()
			defer func() { <-sem }()
			result := processGroup(ctx, g, clientFactory, reservationCmd, notifier)
			mu.Lock()
			results = append(results, result)
			mu.Unlock()
		}(group)
	}
	wg.Wait()

	total, succeeded, failed := 0, 0, 0
	for _, r := range results {
		succeeded += r.succeeded()
		failed += r.failed()
		total += len(r.Items)
	}
	slog.Info("tcm-scheduler job completed",
		slog.Int("total", total),
		slog.Int("succeeded", succeeded),
		slog.Int("failed", failed),
	)

	notifier.Notify(buildSummaryMessage(targetDate.String(), results))

	return nil
}

// buildSummaryMessage は全グループの処理結果から1件の通知メッセージを生成する
func buildSummaryMessage(date string, results []GroupResult) NotifyMessage {
	totalSucceeded, totalFailed := 0, 0
	for _, r := range results {
		totalSucceeded += r.succeeded()
		totalFailed += r.failed()
	}

	level := NotifyLevelInfo

	fields := make([]NotifyField, 0, len(results)+1)
	fields = append(fields, NotifyField{
		Name:  "対象日",
		Value: date,
	})

	for _, r := range results {
		fields = append(fields, NotifyField{
			Name:  fmt.Sprintf("👤 %s", r.DisplayID),
			Value: formatGroupResult(r),
		})
	}

	return NotifyMessage{
		Level:  level,
		Title:  fmt.Sprintf("バッチ処理完了（成功 %d / 失敗 %d）", totalSucceeded, totalFailed),
		Fields: fields,
	}
}

func formatGroupResult(r GroupResult) string {
	if r.MissingCreds {
		return "❌ 認証情報が未設定"
	}
	if r.LoginErr != nil {
		return fmt.Sprintf("❌ ログイン失敗: %s", r.LoginErr.Error())
	}

	lines := make([]string, 0, len(r.Items))
	for _, item := range r.Items {
		timeRange := fmt.Sprintf("%d:%02d〜%d:%02d", item.FromHour, item.FromMinute, item.ToHour, item.ToMinute)
		if item.Success {
			lines = append(lines, fmt.Sprintf("✅ %s | %s", item.RoomID, timeRange))
		} else {
			errMsg := ""
			if item.Err != nil {
				errMsg = fmt.Sprintf(" (%s)", item.Err.Error())
			}
			lines = append(lines, fmt.Sprintf("❌ %s | %s%s", item.RoomID, timeRange, errMsg))
		}
	}
	return strings.Join(lines, "\n")
}

func processGroup(ctx context.Context, group *Group, clientFactory OfficialClientFactory, rsvCmd gateway.ReservationCommand, notifier Notifier) GroupResult {
	master := group.MasterUser

	displayID := master.ID.String()
	if master.OfficialSiteID != nil {
		displayID = *master.OfficialSiteID
	}

	if master.OfficialSiteID == nil || master.OfficialSitePassword == nil {
		slog.Error("master user missing official site credentials",
			slog.String("user_id", master.ID.String()),
		)
		items := make([]ReservationItemResult, len(group.Reservations))
		for i, rsv := range group.Reservations {
			markFailed(ctx, rsvCmd, rsv.Reservation.ID, "missing credentials")
			items[i] = ReservationItemResult{
				RoomID:     rsv.Reservation.RoomID,
				Date:       rsv.Reservation.Date.String(),
				FromHour:   rsv.Reservation.FromHour,
				FromMinute: rsv.Reservation.FromMinute,
				ToHour:     rsv.Reservation.ToHour,
				ToMinute:   rsv.Reservation.ToMinute,
				Success:    false,
			}
		}
		return GroupResult{DisplayID: displayID, MissingCreds: true, Items: items}
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
		items := make([]ReservationItemResult, len(group.Reservations))
		for i, rsv := range group.Reservations {
			markFailed(ctx, rsvCmd, rsv.Reservation.ID, err.Error())
			items[i] = ReservationItemResult{
				RoomID:     rsv.Reservation.RoomID,
				Date:       rsv.Reservation.Date.String(),
				FromHour:   rsv.Reservation.FromHour,
				FromMinute: rsv.Reservation.FromMinute,
				ToHour:     rsv.Reservation.ToHour,
				ToMinute:   rsv.Reservation.ToMinute,
				Success:    false,
				Err:        err,
			}
		}
		return GroupResult{DisplayID: displayID, LoginErr: err, Items: items}
	}

	slog.Info("logged in to official site",
		slog.String("official_site_id", *master.OfficialSiteID),
		slog.Int("reservations_count", len(group.Reservations)),
	)
	notifier.Notify(NotifyMessage{
		Level: NotifyLevelDebug,
		Title: "公式サイトにログイン成功",
		Fields: []NotifyField{
			{Name: "公式サイトID", Value: *master.OfficialSiteID},
			{Name: "予約件数", Value: fmt.Sprintf("%d", len(group.Reservations))},
		},
	})

	// 各予約を公式サイトに投入
	items := make([]ReservationItemResult, 0, len(group.Reservations))
	for _, rsvView := range group.Reservations {
		rsv := rsvView.Reservation

		date := rsv.Date
		tcmDate := tcmrsv.NewDate(date.Year, date.Month, date.Day)

		item := ReservationItemResult{
			RoomID:     rsv.RoomID,
			Date:       rsv.Date.String(),
			FromHour:   rsv.FromHour,
			FromMinute: rsv.FromMinute,
			ToHour:     rsv.ToHour,
			ToMinute:   rsv.ToMinute,
		}

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
			item.Success = false
			item.Err = err
			notifier.Notify(NotifyMessage{
				Level: NotifyLevelDebug,
				Title: "予約失敗（詳細）",
				Fields: []NotifyField{
					{Name: "部屋ID", Value: rsv.RoomID},
					{Name: "日付", Value: rsv.Date.String()},
					{Name: "エラー", Value: err.Error()},
				},
			})
			markFailed(ctx, rsvCmd, rsv.ID, err.Error())
		} else {
			slog.Info("reservation succeeded",
				slog.String("reservation_id", rsv.ID.String()),
				slog.String("room_id", rsv.RoomID),
				slog.String("date", rsv.Date.String()),
			)
			item.Success = true
			notifier.Notify(NotifyMessage{
				Level: NotifyLevelDebug,
				Title: "予約成功（詳細）",
				Fields: []NotifyField{
					{Name: "部屋ID", Value: rsv.RoomID},
					{Name: "日付", Value: rsv.Date.String()},
				},
			})
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
		items = append(items, item)
	}

	return GroupResult{DisplayID: displayID, Items: items}
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
