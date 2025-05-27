package tests

import (
	"os"
	"testing"

	"github.com/ekkx/tcmrsv-web/server/tests/testhelper"
)

func TestMain(m *testing.M) {
	_ = testhelper.GetPool(nil)

	code := m.Run()

	testhelper.ClosePool()

	os.Exit(code)
}
