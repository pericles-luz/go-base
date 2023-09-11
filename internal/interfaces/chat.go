package interfaces

import "errors"

var (
	ErrChatMessageTextEmpty         = errors.New("texto inválido")
	ErrChatMessageNumberEmpty       = errors.New("número inválido")
	ErrChatMessageContactIDEmpty    = errors.New("contactId inválido")
	ErrChatMessageServiceIDEmpty    = errors.New("serviceId inválido")
	ErrChatMessageFileBase64Empty   = errors.New("base64 inválido")
	ErrChatMessageFileMimetypeEmpty = errors.New("mimetype inválido")
	ErrChatMessageFileNameEmpty     = errors.New("name inválido")
	ErrAuthenticationFailed         = errors.New("autenticação falhou")
)

type IChat interface {
	SendMessage(IChatMessage) error
}

type IChatMessage interface {
	GetText() string
	GetNumber() string
	GetContactID() string
	GetServiceID() string
	GetFile() IChatFile
	SetText(string) error
	SetNumber(string) error
	SetContactID(string) error
	SetServiceID(string) error
	SetFile(IChatFile) error
	ToJSON() (string, error)
	Validate() error
}

type IChatFile interface {
	GetBase64() string
	GetMimetype() string
	GetName() string
	SetBase64(string) error
	SetMimetype(string) error
	SetName(string) error
	Validate() error
}
