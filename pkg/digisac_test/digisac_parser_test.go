package digisac_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/digisac"
	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestParser_GetContactMustParseData(t *testing.T) {
	parser := &digisac.Parser{}
	contact, err := parser.GetContact(dataContact())
	require.NoError(t, err)
	require.NotNil(t, contact)
	require.Equal(t, "553186058910", contact.Data.Number)
	require.Equal(t, "31986058910", contact.Phonenumber())
	require.True(t, utils.ValidateUUID(contact.ContactID()))
	t.Log(contact)
}
