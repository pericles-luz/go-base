package digisac_test

import (
	"os"
	"testing"

	"github.com/pericles-luz/go-base/internals/factory"
	"github.com/stretchr/testify/require"
)

func TestDigisac_MustGetContactByIDFromAPI(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	digisac, err := factory.NewDigisac("digisac.sindireceita")
	require.NoError(t, err)
	contact, err := digisac.GetContactByID("ce4af5d4-afd5-4c4f-b5ab-a9ade3c223b0")
	require.NoError(t, err)
	require.NotNil(t, contact)
	require.Equal(t, "553186058910", contact.Data.Number)
	t.Log(contact)
}
