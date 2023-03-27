package d360_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/pericles-luz/go-base/internals/factory"
	"github.com/pericles-luz/go-base/pkg/d360"
	"github.com/stretchr/testify/require"
)

func TestChatD360_Autenticate(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	chatD360, err := factory.NewChatD360("d360.dev")
	require.NoError(t, err)
	token, err := chatD360.Autenticate()
	require.NoError(t, err)
	require.NotNil(t, token)
	require.True(t, token.IsValid())
}

func TestChatD360_SendMessage(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	chatD360, err := factory.NewChatD360("d360.dev")
	require.NoError(t, err)
	var sendMessageRequest d360.D360_MessageRequest
	require.NoError(t, json.Unmarshal([]byte(dataChatD360MessageRequest()), &sendMessageRequest))
	sent, err := chatD360.SendMessage(map[string]interface{}{
		"DE_Telefone": sendMessageRequest.To,
		"TX_Mensagem": sendMessageRequest.Text.Body,
	})
	require.NoError(t, err)
	require.NotEmpty(t, sent)
}

func TestChatD360_SendMessageInteractive(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	chatD360, err := factory.NewChatD360("d360.dev")
	require.NoError(t, err)
	sent, err := chatD360.SendMessageInteractive(dataChatD360InteractiveMessageMap())
	require.NoError(t, err)
	require.NotEmpty(t, sent)
}

func TestChatD360_SendMessageInteractiveWithImage(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	chatD360, err := factory.NewChatD360("d360.dev")
	require.NoError(t, err)
	sent, err := chatD360.SendMessageInteractive(dataChatD360InteractiveMessageWithImageMap())
	require.NoError(t, err)
	require.NotEmpty(t, sent)
}

func TestChatD360_GetTemplateInteractiveResponse(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	// só funciona em produção
	chatD360, err := factory.NewChatD360("d360.prod")
	require.NoError(t, err)
	got, err := chatD360.GetTemplateInteractive()
	require.NoError(t, err)
	require.NotEmpty(t, got)
}

func TestChatD360_SendTemplateMessage(t *testing.T) {
	if os.Getenv("GITHUB") == "yes" {
		t.Skip("Não testar no github")
	}
	// só funciona em produção
	chatD360, err := factory.NewChatD360("d360.prod")
	require.NoError(t, err)
	got, err := chatD360.SendMessageTemplate(dataChatD360InteractiveTemplateWithImageMap())
	require.NoError(t, err)
	require.NotEmpty(t, got)
}
