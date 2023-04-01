package digisac_test

import (
	"encoding/json"
	"testing"

	"github.com/pericles-luz/go-base/pkg/digisac"
	"github.com/stretchr/testify/require"
)

func TestDigisac_UnMarshalContactToContact(t *testing.T) {
	var contact digisac.Contact
	require.NoError(t, json.Unmarshal([]byte(dataContact()), &contact))
	require.Equal(t, "553186058910", contact.Data.Number)
	t.Log(contact)
}
