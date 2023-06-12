package conf_test

import (
	"os"
	"testing"

	"github.com/pericles-luz/go-base/pkg/conf"
	"github.com/stretchr/testify/require"
)

func TestMailer(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Skip when running on github")
	}
	mailer := conf.NewMailer()
	require.NoError(t, mailer.Load("mailer.test"))
	require.Equal(t, "smtp.test.email", mailer.GetConfig()["host"])
	require.Equal(t, "995", mailer.GetConfig()["port"])
	require.Equal(t, "test@opssbank.com.br", mailer.GetConfig()["username"])
	require.Equal(t, "SupersenhaTeste", mailer.GetConfig()["password"])
}
