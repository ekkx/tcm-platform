package testhelper

import (
	"context"
	"testing"

	"github.com/ekkx/tcmrsv-web/server/internal/config"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/stretchr/testify/require"
)

func GetContextWithConfig(t *testing.T) context.Context {
	t.Helper()

	cfg, err := config.New()
	require.NoError(t, err)

	return ctxhelper.SetConfig(context.Background(), cfg)
}
