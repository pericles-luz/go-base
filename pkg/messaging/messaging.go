package messaging

type Channels struct {
	// Channel for sending messages to the messaging adapter
	SMS chan map[string]string
	// Channel for sending messages to the messaging adapter
	Email chan map[string]interface{}
}

func NewChannels() *Channels {
	return &Channels{
		SMS:   make(chan map[string]string, 100),
		Email: make(chan map[string]interface{}, 100),
	}
}
