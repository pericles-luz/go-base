package digisac

import (
	"time"

	"github.com/pericles-luz/go-base/pkg/utils"
)

type Contact struct {
	Unsubscribed         bool        `json:"unsubscribed,omitempty"`
	ID                   string      `json:"id,omitempty"`
	IsMe                 bool        `json:"isMe,omitempty"`
	IsGroup              bool        `json:"isGroup,omitempty"`
	IsBroadcast          bool        `json:"isBroadcast,omitempty"`
	Unread               int         `json:"unread,omitempty"`
	IsSilenced           bool        `json:"isSilenced,omitempty"`
	IsMyContact          bool        `json:"isMyContact,omitempty"`
	HadChat              bool        `json:"hadChat,omitempty"`
	Visible              bool        `json:"visible,omitempty"`
	Note                 string      `json:"note,omitempty"`
	LastMessageAt        time.Time   `json:"lastMessageAt,omitempty"`
	LastMessageID        string      `json:"lastMessageId,omitempty"`
	AccountID            string      `json:"accountId,omitempty"`
	ServiceID            string      `json:"serviceId,omitempty"`
	PersonID             string      `json:"personId,omitempty"`
	DefaultDepartmentID  string      `json:"defaultDepartmentId,omitempty"`
	DefaultUserID        time.Time   `json:"defaultUserId,omitempty"`
	CreatedAt            time.Time   `json:"createdAt,omitempty"`
	UpdatedAt            time.Time   `json:"updatedAt,omitempty"`
	DeletedAt            time.Time   `json:"deletedAt,omitempty"`
	Status               string      `json:"status,omitempty"`
	LastContactMessageAt time.Time   `json:"lastContactMessageAt,omitempty"`
	Name                 string      `json:"name,omitempty"`
	InternalName         string      `json:"internalName,omitempty"`
	AlternativeName      string      `json:"alternativeName,omitempty"`
	Data                 ContactData `json:"data,omitempty"`
	LastMessage          time.Time   `json:"lastMessage,omitempty"`
}
type Webchat struct {
}
type ContactData struct {
	Valid              bool      `json:"valid,omitempty"`
	Number             string    `json:"number,omitempty"`
	Unread             bool      `json:"unread,omitempty"`
	Webchat            Webchat   `json:"webchat,omitempty"`
	IsOriginal         bool      `json:"isOriginal,omitempty"`
	BotIsRunning       bool      `json:"botIsRunning,omitempty"`
	LastChargedMessage time.Time `json:"lastChargedMessage,omitempty"`
}

func (c *Contact) Phonenumber() string {
	if c.Data.Number == "" {
		return ""
	}
	if len(c.Data.Number) == utils.FULL_PHONENUMBER_LENGTH {
		return c.Data.Number[2:]
	}
	return utils.WhatsappNumberToBrazilianPhonenumber(c.Data.Number)
}

func (c *Contact) ContactID() string {
	return c.ID
}
