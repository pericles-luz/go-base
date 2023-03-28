package conf_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/conf"
	"github.com/stretchr/testify/require"
)

func TestDatabaseMustLoadJSONCorrectly(t *testing.T) {
	conn, err := conf.NewDatabase("agnu.dev")
	require.NoError(t, err)
	require.NotNil(t, conn)
}
