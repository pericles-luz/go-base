package conf_test

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/pericles-luz/go-base/pkg/conf"
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
	require.NotNil(t, cfg.DbConfiguration)
	err = cfg.DbConfiguration.Validate()
	require.NoError(t, err)
	dsn, err := cfg.DbConfiguration.GetDSN()
	require.NoError(t, err)
	require.NotNil(t, dsn)
	require.Equal(t, fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		cfg.DbConfiguration.Username,
		url.QueryEscape(cfg.DbConfiguration.Password),
		cfg.DbConfiguration.Host,
		cfg.DbConfiguration.Port,
		cfg.DbConfiguration.DBName,
	), dsn)
}

func TestInitialConfigurationPostgres(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Skip when running on github")
	}
	cfg, err := conf.NewInitialConfig("initial_test_pg")
	require.NoError(t, err)
	err = cfg.Validate()
	require.NoError(t, err)
	require.NotNil(t, cfg.DbConfiguration)
	err = cfg.DbConfiguration.Validate()
	require.NoError(t, err)
	dsn, err := cfg.DbConfiguration.GetDSN()
	require.NoError(t, err)
	require.NotNil(t, dsn)
	require.Equal(t, fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		cfg.DbConfiguration.Username,
		url.QueryEscape(cfg.DbConfiguration.Password),
		cfg.DbConfiguration.Host,
		cfg.DbConfiguration.Port,
		cfg.DbConfiguration.DBName,
	), dsn)
}

func TestInitialConfigurationFailIfFileDoesNotExists(t *testing.T) {
	_, err := conf.NewInitialConfig("initial_test_fail")
	require.Error(t, err)
}
