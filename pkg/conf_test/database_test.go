package conf_test

import (
	"os"
	"testing"

	"github.com/pericles-luz/go-base/pkg/conf"
	"github.com/stretchr/testify/require"
)

func TestDatabaseMustLoadJSONCorrectly(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Skip when running on github")
	}
	conn, err := conf.NewDatabase("agnu.dev")
	require.NoError(t, err)
	require.NotNil(t, conn)
}
