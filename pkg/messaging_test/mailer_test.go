package messaging_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/conf"
	"github.com/pericles-luz/go-base/pkg/messaging"
	"github.com/stretchr/testify/require"
)

func TestMailer_Send(t *testing.T) {
	t.Skip("Test only if necessary")
	mailer := messaging.NewMailer()
	mailer.SetTo("pericles.luz@gmail.com")
	mailer.SetSubject("Teste via golang")
	mailer.SetBody(`<!DOCTYPE html>
	<html>
	<body>
			<h3>Name:</h3><span>{{.Name}}</span><br/><br/>
			<h3>Email:</h3><span>{{.Message}}</span><br/>
	</body>
	</html>`)
	config := conf.NewMailer()
	require.NoError(t, config.Load("mailer.prod"))
	mailer.SetConfig(config.GetConfig())
	require.NoError(t, mailer.Send())
}
