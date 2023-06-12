package conf_test

import (
	"os"
	"testing"

	"github.com/pericles-luz/go-base/pkg/conf"
	"github.com/stretchr/testify/require"
)

func TestDigisac_MustLoagConfigFile(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Skip when running on github")
	}
	cfg := conf.NewDigisac()
	err := cfg.Load("digisac.dev")
	require.NoError(t, err)
	require.Equal(t, "https://dev.digisac.chat/api/v1", cfg.LinkAPI)
	require.NotEmpty(t, cfg.Token)
}
