/*
tcmscheduler は、東京音楽大学の公式サイトへの予約を自動化するためのスケジューラーです。

仕様:
- 毎日 12:00 に実行される
- 当日の予約一覧を取得する
- 予約したユーザーのマスターユーザーごとに予約をグループ化する（マスターユーザーはそのアカウントを使用）
- グループ化した予約をマスターユーザーごとに処理する
- 公式サイトへの予約処理は非同期で行う
- 予約処理の結果はまとめる
- それぞれの予約が成功したか失敗したかは、内部のデータベースに記録する
- 失敗した予約は、再度リトライするか、discord などに監査通知する
*/
package main

import (
	"log"
	"log/slog"
	"time"

	"github.com/ekkx/tcmrsv-web/internal/config"
	"github.com/ekkx/tcmrsv-web/internal/shared/logger"
	"github.com/robfig/cron/v3"
)

func jst() *time.Location {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalf("failed to load location: %v", err)
	}
	return loc
}

func main() {
	// This is the entry point for the TCMScheduler application.
	// The main function will initialize the application and start the scheduler.

	// Initialize the application components here
	// initializeApp()

	// Start the scheduler
	// startScheduler()

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger.Init(cfg)

	slog.Info("TCMScheduler started")

	slog.Debug("Configuration loaded", slog.String("cron_expression", cfg.Scheduler.CronExpression))

	c := cron.New(
		cron.WithLocation(jst()),
		cron.WithSeconds(),
	)

	c.AddFunc(cfg.Scheduler.CronExpression, func() {
		if err := RunBatch(); err != nil {
			slog.Error("batch error: %v", err.Error(), err)
		}
	})

	c.Start()
	select {}
}

func RunBatch() error {
	slog.Info("Running batch job...")
	return nil
}
