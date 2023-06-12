package conf_test

import (
	"os"
	"testing"

	"github.com/pericles-luz/go-base/pkg/conf"
	"github.com/stretchr/testify/require"
)

func TestNewWhatsappTemplate(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Skip when running on github")
	}
	cfg := conf.NewWhatsappTemplate()
	err := cfg.Load("whatsapp.template.test")
	require.NoError(t, err)
	require.Equal(t, "test", cfg.Name)
	require.Equal(t, "39721bde_f26f_42f3_b928_aa4267759d7c", cfg.Namespace)
}
