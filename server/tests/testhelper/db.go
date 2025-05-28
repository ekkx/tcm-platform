package testhelper

import (
	"context"
	"sync"
	"testing"

	"github.com/ekkx/tcmrsv-web/server/internal/config"
	"github.com/ekkx/tcmrsv-web/server/pkg/database"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

var (
	pool     *pgxpool.Pool
	initPool sync.Once
)

func GetPool(t *testing.T) *pgxpool.Pool {
	t.Helper()

	initPool.Do(func() {
		cfg, err := config.New()
		require.NoError(t, err)

		pool, err = cfg.Database.Open()
		require.NoError(t, err)
	})

	return pool
}

func ClosePool() {
	if pool != nil {
		pool.Close()
	}
}

// トランザクションのロールバックを利用することによってクリーンアップを自動化する
func RunWithTx(t *testing.T, testFn func(db database.Execer)) {
	t.Helper()

	ctx := context.Background()
	tx, err := GetPool(t).Begin(ctx)
	require.NoError(t, err)

	defer tx.Rollback(ctx)

	testFn(tx)
}
