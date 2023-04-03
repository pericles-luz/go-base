package digisac_test

import (
	"encoding/json"
	"testing"

	"github.com/pericles-luz/go-base/pkg/digisac"
	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestDigisac_UnMarshalContactToContact(t *testing.T) {
	var contact digisac.Contact
	require.NoError(t, json.Unmarshal([]byte(dataContact()), &contact))
	require.Equal(t, "553186058910", contact.Data.Number)
	require.Equal(t, "31986058910", contact.Phonenumber())
	require.True(t, utils.ValidateUUID(contact.ContactID()))
	t.Log(contact)
}

func TestDigisac_UnMarshalContactToContactNineBaseDigits(t *testing.T) {
	var contact digisac.Contact
	require.NoError(t, json.Unmarshal([]byte(dataContactNineBaseDigits()), &contact))
	require.Equal(t, "5521983027896", contact.Data.Number)
	require.Equal(t, "5521983027896", contact.Phonenumber())
	require.True(t, utils.ValidateUUID(contact.ContactID()))
	t.Log(contact)
}
