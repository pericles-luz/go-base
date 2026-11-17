package d360_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pericles-luz/go-base/pkg/d360"
	"github.com/pericles-luz/go-rest/pkg/rest"
	"github.com/stretchr/testify/require"
)

func newD360WithServer(t *testing.T, status int, body string) *d360.D360 {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)
	return d360.NewChatD360(rest.NewRest(map[string]interface{}{
		"token":   "chave-de-teste",
		"linkAPI": server.URL,
	}))
}

func templateForHTTPTest() map[string]interface{} {
	return d360.NewChatD360TemplateTextMessage(map[string]interface{}{
		"DE_Telefone":  "61999999999",
		"DE_Namespace": "namespace",
		"DE_Nome":      "alerta",
		"parametros": []map[string]interface{}{
			{"DE_Tipo": "text", "DE_Texto": "mensagem"},
		},
	})
}

func TestChatD360_SendMessageTemplateMustFailWhenAPIReturnsHTTPError(t *testing.T) {
	chatD360 := newD360WithServer(t, http.StatusBadRequest,
		`{"error":[{"code":132018,"title":"There was an issue with the parameters in your template."}]}`)
	sent, err := chatD360.SendMessageTemplate(templateForHTTPTest())
	require.Error(t, err)
	require.Nil(t, sent)
	require.Contains(t, err.Error(), "HTTP 400")
	require.Contains(t, err.Error(), "132018")
}

func TestChatD360_SendMessageTemplateMustFailWhenRateLimited(t *testing.T) {
	chatD360 := newD360WithServer(t, http.StatusTooManyRequests, `{"error":[{"code":130429}]}`)
	_, err := chatD360.SendMessageTemplate(templateForHTTPTest())
	require.ErrorContains(t, err, "HTTP 429")
}

func TestChatD360_SendMessageTemplateMustSucceedWhenAPIAcceptsMessage(t *testing.T) {
	chatD360 := newD360WithServer(t, http.StatusOK,
		`{"messaging_product":"whatsapp","contacts":[{"input":"5561999999999","wa_id":"5561999999999"}],"messages":[{"id":"wamid.teste"}]}`)
	sent, err := chatD360.SendMessageTemplate(templateForHTTPTest())
	require.NoError(t, err)
	require.Len(t, sent, 1)
	require.Equal(t, "wamid.teste", sent[0].GetID())
}
