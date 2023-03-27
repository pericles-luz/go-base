package conf_test

import (
	"os"
	"testing"

	"github.com/pericles-luz/go-base/pkg/conf"
	"github.com/stretchr/testify/require"
)

func TestChatD360Config(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	cfg := conf.NewChatD360Config()
	err := cfg.Load("d360.dev")
	require.NoError(t, err)
	require.Equal(t, "https://waba-sandbox.360dialog.io/v1", cfg.LinkAPI)
}
