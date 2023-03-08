package conf_test

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/pericles-luz/go-base/pkg/infra/conf"
	"github.com/stretchr/testify/require"
)

func TestInitialConfiguration(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Skip when running on github")
	}
	cfg, err := conf.NewInitialConfig("initial")
	require.NoError(t, err)
	err = cfg.Validate()
	require.NoError(t, err)
	require.NotNil(t, cfg.DBConfiguration)
	err = cfg.DBConfiguration.Validate()
	require.NoError(t, err)
	dsn, err := cfg.DBConfiguration.GetDSN()
	require.NoError(t, err)
	require.NotNil(t, dsn)
	require.Equal(t, fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		cfg.DBConfiguration.Username,
		url.QueryEscape(cfg.DBConfiguration.Password),
		cfg.DBConfiguration.Host,
		cfg.DBConfiguration.Port,
		cfg.DBConfiguration.DBName,
	), dsn)
}

func TestInitialConfigurationFailIfFileDoesNotExists(t *testing.T) {
	_, err := conf.NewInitialConfig("initial_test_fail")
	require.Error(t, err)
}
