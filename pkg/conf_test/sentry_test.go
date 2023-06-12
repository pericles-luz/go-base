package conf_test

import (
	"os"
	"testing"

	"github.com/pericles-luz/go-base/pkg/conf"
	"github.com/stretchr/testify/require"
)

func TestSentry(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Skip when running on github")
	}
	sentry := conf.NewSentry()
	require.NoError(t, sentry.Load("sentry"))
	require.Greater(t, len(sentry.GetConfig()["dsn"]), 50)
}
