package d360_test

import (
	"encoding/json"
	"testing"

	"github.com/pericles-luz/go-base/internals/factory"
	"github.com/pericles-luz/go-base/pkg/d360"
	"github.com/stretchr/testify/require"
)

func TestChatD360_Autenticate(t *testing.T) {
	t.Skip("Test only if necessary")
	chatD360, err := factory.NewChatD360("d360.dev")
	require.NoError(t, err)
	token, err := chatD360.Autenticate()
	require.NoError(t, err)
	require.NotNil(t, token)
	require.True(t, token.IsValid())
}

func TestChatD360_SendMessage(t *testing.T) {
	t.Skip("Test only if necessary")
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
	t.Skip("Test only if necessary")
	chatD360, err := factory.NewChatD360("d360.dev")
	require.NoError(t, err)
	sent, err := chatD360.SendMessageInteractive(dataChatD360InteractiveMessageMap())
	require.NoError(t, err)
	require.NotEmpty(t, sent)
}

func TestChatD360_SendMessageInteractiveWithImage(t *testing.T) {
	t.Skip("Test only if necessary")
	chatD360, err := factory.NewChatD360("d360.dev")
	require.NoError(t, err)
	sent, err := chatD360.SendMessageInteractive(dataChatD360InteractiveMessageWithImageMap())
	require.NoError(t, err)
	require.NotEmpty(t, sent)
}

func TestChatD360_SendMessageInteractiveWithPDF(t *testing.T) {
	t.Skip("Test only if necessary")
	chatD360, err := factory.NewChatD360("d360.sindireceita")
	require.NoError(t, err)
	sent, err := chatD360.SendMessageInteractive(dataChatD360InteractiveMessageWithPDFMap())
	require.NoError(t, err)
	require.NotEmpty(t, sent)
}

func TestChatD360_GetTemplateInteractiveResponse(t *testing.T) {
	t.Skip("Test only if necessary")
	// só funciona em produção
	chatD360, err := factory.NewChatD360("d360.prod")
	require.NoError(t, err)
	got, err := chatD360.GetTemplateInteractive()
	t.Log("got", got)
	require.NoError(t, err)
	require.NotEmpty(t, got)
}

func TestChatD360_SendTemplateMessage(t *testing.T) {
	t.Skip("Test only if necessary")
	// só funciona em produção
	chatD360, err := factory.NewChatD360("d360.prod")
	require.NoError(t, err)
	got, err := chatD360.SendMessageTemplate(dataChatD360InteractiveTemplateWithImageMap())
	require.NoError(t, err)
	require.NotEmpty(t, got)
}

func TestChatD360_SendTextTemplateMessageToken(t *testing.T) {
	t.Skip("Test only if necessary")
	// só funciona em produção
	chatD360, err := factory.NewChatD360("d360.prod")
	require.NoError(t, err)
	got, err := chatD360.SendMessageTemplate(dataChatD360TextTemplateMessageMap())
	require.NoError(t, err)
	require.NotEmpty(t, got)
}

func TestChatD360_SendTextTemplateMessageOla(t *testing.T) {
	t.Skip("Test only if necessary")
	// só funciona em produção
	chatD360, err := factory.NewChatD360("d360.oopss")
	require.NoError(t, err)
	got, err := chatD360.SendMessageTemplate(dataChatD360TextTemplateMessageOlaMap())
	require.NoError(t, err)
	require.NotEmpty(t, got)
}

func TestChatD360_SendTextTemplateSignatureAndCertificateMessage(t *testing.T) {
	t.Skip("Test only if necessary")
	// só funciona em produção
	chatD360, err := factory.NewChatD360("d360.oopss")
	require.NoError(t, err)
	got, err := chatD360.SendMessageTemplate(dataChatD360TextTemplateMessageLinkMap())
	require.NoError(t, err)
	require.NotEmpty(t, got)
}
