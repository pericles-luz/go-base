package utils

const (
	TEST_CPF   = "00000000191"
	TEST_TOKEN = "123456"

	// formas de envio de mensagem
	SEND_MEDIA_SMS      = 1
	SEND_MEDIA_EMAIL    = 2
	SEND_MEDIA_WHATSAPP = 3

	PHONENUMBER_MAX_LENGTH = 9
	PHONENUMBER_MIN_LENGTH = 8

	WHATSAPP_PHONENUMBER_LENGTH = 12

	VALID_EMAIL_REGEX = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)
