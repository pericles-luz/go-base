package utils

const (
	TEST_CPF                 = "00000000191"
	TEST_NON_MEMBER_CPF      = "00000000272"
	TEST_WITHOUT_ANSWERS_CPF = "00000000353"
	TEST_TOKEN               = "123456"

	// formas de envio de mensagem
	SEND_MEDIA_SMS   = 1
	SEND_MEDIA_EMAIL = 2

	PHONENUMBER_MAX_LENGTH = 9
	PHONENUMBER_MIN_LENGTH = 8

	VALID_EMAIL_REGEX = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)
