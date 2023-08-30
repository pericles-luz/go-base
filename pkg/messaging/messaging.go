package messaging

type Channels struct {
	// Channel for sending messages to the messaging adapter
	// Each message must have the following structure:
	// {
	// 	"EM_Destinatario": "email address",
	// 	"DE_Body": "message body",
	//  "DE_Assunto": "message subject
	// }
	Email chan map[string]interface{}
	// Channel for sending messages to the messaging adapter
	// Each message must have the following structure:
	// {
	// 	"DE_Telefone": "phone number",
	// 	"DE_Mensagem": "message body",
	// }
	SMS chan map[string]string
	// Channel for sending messages to the messaging adapter
	// Each message must have the following structure:
	// {
	// 	"DE_Telefone": "phone number",
	// 	"DE_Mensagem": "message body",
	// }
	Whatsapp chan map[string]interface{}
	// Channel for sending messages to the messaging adapter
	// Each message must have the following structure:
	// {
	// 	"exchange": "exchange name",
	// 	"routingKey": "routing key",
	// 	"body": "message body",
	// 	"durable": "if the message is durable",
	// }
	Messaging chan map[string]interface{}
	// Channel for sending shutdown signal to the web server
	Shutdown chan bool
}

func NewChannels() *Channels {
	return &Channels{
		SMS:       make(chan map[string]string, 100),
		Email:     make(chan map[string]interface{}, 100),
		Whatsapp:  make(chan map[string]interface{}, 100),
		Messaging: make(chan map[string]interface{}, 100),
	}
}
