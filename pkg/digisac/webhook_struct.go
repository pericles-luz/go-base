package digisac

import "time"

const (
	MESSAGE_TYPE_TICKET   = "ticket"
	MESSAGE_TYPE_CHAT     = "chat"
	MESSAGE_ORIGIN_TICKET = "ticket"
	EVENT_NEW_MESSAGE     = "message.created"
)

type WebHookMessage struct {
	Event     string      `json:"event,omitempty"`
	Data      WebhookData `json:"data,omitempty"`
	WebhookID string      `json:"webhookId,omitempty"`
	Timestamp time.Time   `json:"timestamp,omitempty"`
}
type InternalData struct {
	Ack            int  `json:"ack,omitempty"`
	IsNew          bool `json:"isNew,omitempty"`
	IsFirst        bool `json:"isFirst,omitempty"`
	TicketTransfer bool `json:"ticketTransfer,omitempty"`
	DontOpenTicket bool `json:"dontOpenTicket,omitempty"`
	TicketClose    bool `json:"ticketClose,omitempty"`
}
type WebhookData struct {
	ID                 string                 `json:"id,omitempty"`
	IsFromMe           bool                   `json:"isFromMe,omitempty"`
	Sent               bool                   `json:"sent,omitempty"`
	Type               string                 `json:"type,omitempty"`
	Timestamp          time.Time              `json:"timestamp,omitempty"`
	Data               InternalData           `json:"data,omitempty"`
	Visible            bool                   `json:"visible,omitempty"`
	AccountID          string                 `json:"accountId,omitempty"`
	ContactID          string                 `json:"contactId,omitempty"`
	FromID             string                 `json:"fromId,omitempty"`
	ServiceID          string                 `json:"serviceId,omitempty"`
	ToID               string                 `json:"toId,omitempty"`
	UserID             string                 `json:"userId,omitempty"`
	TicketID           string                 `json:"ticketId,omitempty"`
	TicketUserID       string                 `json:"ticketUserId,omitempty"`
	TicketDepartmentID string                 `json:"ticketDepartmentId,omitempty"`
	QuotedMessageID    string                 `json:"quotedMessageId,omitempty"`
	Origin             string                 `json:"origin,omitempty"`
	HsmID              string                 `json:"hsmId,omitempty"`
	Text               string                 `json:"text,omitempty"`
	Obfuscated         bool                   `json:"obfuscated,omitempty"`
	Files              map[string]interface{} `json:"files,omitempty"`
	QuotedMessage      string                 `json:"quotedMessage,omitempty"`
	IsFromBot          bool                   `json:"isFromBot,omitempty"`
}

func (w *WebHookMessage) IsNewMessage() bool {
	return w.Event == EVENT_NEW_MESSAGE
}

func (w *WebHookMessage) IsFromContact() bool {
	if w.Data.Type != MESSAGE_TYPE_CHAT {
		return false
	}
	if w.Data.IsFromMe || w.Data.IsFromBot {
		return false
	}
	return true
}

func (w *WebHookMessage) ContactID() string {
	return w.Data.ContactID
}
