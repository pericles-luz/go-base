package digisac_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/digisac"
	"github.com/stretchr/testify/require"
)

func TestParser_GetContacrtMustParserJ(t *testing.T) {
	parser := &digisac.Parser{}
	contact, err := parser.GetContact(dataContact())
	require.NoError(t, err)
	require.NotNil(t, contact)
	require.Equal(t, "553186058910", contact.Data.Number)
	t.Log(contact)
}
