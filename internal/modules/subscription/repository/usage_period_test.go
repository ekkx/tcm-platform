package repository_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ekkx/tcm-platform/internal/gen/sqlc"
	"github.com/ekkx/tcm-platform/internal/modules/subscription/repository"
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func setupTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://tcmrsv:tcmrsv@tcmrsv-db:5432/tcmrsv_db?sslmode=disable"
	}
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	t.Cleanup(func() { pool.Close() })
	return pool
}

func int32Ptr(v int32) *int32 { return &v }
func strPtr(v string) *string { return &v }

// TestUsagePeriodReset は課金サイクルベースの利用枠リセットを検証する。
//
// 利用枠は「予約対象日」ではなく「予約作成日時（create_time）」が
// 課金期間内かどうかで判断される。
//
// シナリオ:
//   - ユーザーが4/15に契約開始（period: 4/15〜5/15）
//   - 期間内（create_time=4/20）に6月の部屋を2時間予約 → 120分としてカウント
//   - 期間内（create_time=4/25）に4月の部屋を1.5時間予約 → 90分としてカウント
//   - 期間外（create_time=3/10）に4月の部屋を3時間予約 → カウントされない
//   - 期間を翌月に更新 → 全てリセット
func TestUsagePeriodReset(t *testing.T) {
	pool := setupTestDB(t)
	ctx := context.Background()
	querier := sqlc.New(pool)
	repo := repository.New(querier)

	userID := ulid.New()

	// テストユーザーを作成
	_, err := pool.Exec(ctx,
		"INSERT INTO users (id, password, display_name) VALUES ($1, $2, $3)",
		userID, "testpass", "Test User")
	if err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}
	t.Cleanup(func() {
		pool.Exec(ctx, "DELETE FROM reservations WHERE user_id = $1", userID)
		pool.Exec(ctx, "DELETE FROM subscriptions WHERE user_id = $1", userID)
		pool.Exec(ctx, "DELETE FROM users WHERE id = $1", userID)
	})

	// サブスクリプションを作成（period: 4/15〜5/15）
	periodStart := time.Date(2026, 4, 15, 6, 0, 0, 0, time.UTC)
	periodEnd := time.Date(2026, 5, 15, 6, 0, 0, 0, time.UTC)
	subID, err := repo.CreateSubscription(ctx, &repository.CreateSubscriptionParams{
		UserID:             userID,
		StripeCustomerID:   strPtr("cus_test_usage"),
		Plan:               sqlc.PlanTypeLite,
		MonthlyHours:       int32Ptr(30),
		Status:             "active",
		CurrentPeriodStart: &periodStart,
		CurrentPeriodEnd:   &periodEnd,
	})
	if err != nil {
		t.Fatalf("failed to create subscription: %v", err)
	}

	// 期間内に作成（create_time=4/20）、6月の部屋を予約（10:00〜12:00 = 120分）
	// → 予約対象日は期間外だが、作成日時が期間内なのでカウントされる
	_, err = pool.Exec(ctx,
		`INSERT INTO reservations (id, user_id, campus_type, room_id, date, from_hour, from_minute, to_hour, to_minute, status, create_time)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		ulid.New(), userID, "nakameguro", "room-1", "2026-06-10", 10, 0, 12, 0, "pending",
		time.Date(2026, 4, 20, 10, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatalf("failed to create future-date reservation: %v", err)
	}

	// 期間内に作成（create_time=4/25）、4月の部屋を予約（9:00〜10:30 = 90分）
	_, err = pool.Exec(ctx,
		`INSERT INTO reservations (id, user_id, campus_type, room_id, date, from_hour, from_minute, to_hour, to_minute, status, create_time)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		ulid.New(), userID, "nakameguro", "room-1", "2026-04-28", 9, 0, 10, 30, "pending",
		time.Date(2026, 4, 25, 14, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatalf("failed to create in-period reservation: %v", err)
	}

	// 期間外に作成（create_time=3/10）、4月の部屋を予約（14:00〜17:00 = 180分）
	// → 予約対象日は期間内だが、作成日時が期間外なのでカウントされない
	_, err = pool.Exec(ctx,
		`INSERT INTO reservations (id, user_id, campus_type, room_id, date, from_hour, from_minute, to_hour, to_minute, status, create_time)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		ulid.New(), userID, "nakameguro", "room-1", "2026-04-20", 14, 0, 17, 0, "pending",
		time.Date(2026, 3, 10, 9, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatalf("failed to create out-of-period reservation: %v", err)
	}

	// テスト1: create_timeが期間内の予約のみカウント（120 + 90 = 210分）
	usedMinutes, err := repo.GetUsedMinutesByUserID(ctx, userID)
	if err != nil {
		t.Fatalf("failed to get used minutes: %v", err)
	}
	if usedMinutes != 210 {
		t.Errorf("expected 210 used minutes (120+90), got %d", usedMinutes)
	}

	// テスト2: 期間を翌月に更新（5/15〜6/15）→ 利用枠がリセットされる
	newPeriodStart := time.Date(2026, 5, 15, 6, 0, 0, 0, time.UTC)
	newPeriodEnd := time.Date(2026, 6, 15, 6, 0, 0, 0, time.UTC)
	err = repo.UpdateSubscription(ctx, &repository.UpdateSubscriptionParams{
		ID:                 *subID,
		CurrentPeriodStart: &newPeriodStart,
		CurrentPeriodEnd:   &newPeriodEnd,
	})
	if err != nil {
		t.Fatalf("failed to update subscription period: %v", err)
	}

	usedMinutes, err = repo.GetUsedMinutesByUserID(ctx, userID)
	if err != nil {
		t.Fatalf("failed to get used minutes after period update: %v", err)
	}
	if usedMinutes != 0 {
		t.Errorf("expected 0 used minutes after period reset to 5/15-6/15, got %d", usedMinutes)
	}

	// テスト3: 新しい期間内に予約を追加（create_time=5/20、7月の部屋を予約、9:00〜10:30 = 90分）
	_, err = pool.Exec(ctx,
		`INSERT INTO reservations (id, user_id, campus_type, room_id, date, from_hour, from_minute, to_hour, to_minute, status, create_time)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		ulid.New(), userID, "nakameguro", "room-1", "2026-07-05", 9, 0, 10, 30, "pending",
		time.Date(2026, 5, 20, 11, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatalf("failed to create new period reservation: %v", err)
	}

	usedMinutes, err = repo.GetUsedMinutesByUserID(ctx, userID)
	if err != nil {
		t.Fatalf("failed to get used minutes in new period: %v", err)
	}
	if usedMinutes != 90 {
		t.Errorf("expected 90 used minutes in new period 5/15-6/15, got %d", usedMinutes)
	}
}
