package conf_test

import (
	"os"
	"testing"

	"github.com/pericles-luz/go-base/pkg/conf"
	"github.com/stretchr/testify/require"
)

func TestNewVoiceConfig(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Skip when running on github")
	}
	cfg := conf.NewNewVoiceConfig()
	err := cfg.Load("newvoice.test")
	require.NoError(t, err)
	require.Equal(t, "http://api.nvtelecom.com.br", cfg.LinkAPI)
	require.Equal(t, "code", cfg.Code)
	require.Equal(t, "account", cfg.Account)
}
