package ftp_test

import (
	"os"
	"testing"

	"github.com/pericles-luz/go-base/pkg/ftp"
	"github.com/stretchr/testify/require"
)

func TestConfigMustLoad(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Skip when running on github")
	}
	config := ftp.NewConfig()
	err := config.Load("ftp.discadora")
	require.NoError(t, err)
	require.NotEmpty(t, config.URL)
	require.NotEmpty(t, config.Username)
	require.NotEmpty(t, config.Password)
}
