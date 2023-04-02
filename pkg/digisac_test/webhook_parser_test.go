package digisac_test

import (
	"encoding/json"
	"testing"

	"github.com/pericles-luz/go-base/pkg/digisac"
	"github.com/stretchr/testify/require"
)

func TestDigisac_UnMarshalMessageReceivingToWebhookMessage(t *testing.T) {
	var webhookMessage digisac.WebHookMessage
	require.NoError(t, json.Unmarshal([]byte(dataWebhookMessageCreatedReceiving()), &webhookMessage))
	require.True(t, webhookMessage.IsFromContact())
	require.Equal(t, "irpf", webhookMessage.Text())
	t.Log(webhookMessage)
}

func TestDigisac_UnMarshalMessageTransferingToWebhookMessage(t *testing.T) {
	var webhookMessage digisac.WebHookMessage
	require.NoError(t, json.Unmarshal([]byte(dataWebhookMessageCreatedTransfering()), &webhookMessage))
	require.True(t, webhookMessage.Data.Data.TicketTransfer)
	require.False(t, webhookMessage.Data.IsFromMe)
	t.Log(webhookMessage)
}

func TestDigisac_UnMarshalMessageSendingToWebhookMessage(t *testing.T) {
	var webhookMessage digisac.WebHookMessage
	require.NoError(t, json.Unmarshal([]byte(dataWebhookMessageCreatedSending()), &webhookMessage))
	require.True(t, webhookMessage.Data.IsFromMe)
	require.Equal(t, digisac.MESSAGE_TYPE_CHAT, webhookMessage.Data.Type)
	t.Log(webhookMessage)
}

func TestDigisac_UnMarshalMessageCloseingToWebhookMessage(t *testing.T) {
	var webhookMessage digisac.WebHookMessage
	require.NoError(t, json.Unmarshal([]byte(dataWebhookMessageCreatedClosing()), &webhookMessage))
	require.True(t, webhookMessage.Data.Data.TicketClose)
	require.Equal(t, digisac.MESSAGE_TYPE_TICKET, webhookMessage.Data.Type)
	require.Equal(t, digisac.MESSAGE_ORIGIN_TICKET, webhookMessage.Data.Origin)
	t.Log(webhookMessage)
}
