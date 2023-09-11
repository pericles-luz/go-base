package migration

type MessageService struct {
	persistence *MessageDB
}

func NewMessageService(persistence *MessageDB) *MessageService {
	return &MessageService{persistence: persistence}
}

func (s *MessageService) Create(exchange string, routingKey string, data string, durable int16) (*Message, error) {
	message := NewMessage()
	message.SetExchange(exchange)
	message.SetRoutingKey(routingKey)
	message.SetData(data)
	message.SetDurable(durable)
	err := message.IsValid()
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (s *MessageService) Send(exchange string, routingKey string, data string, durable int16) (*Message, error) {
	message, err := s.Create(exchange, routingKey, data, durable)
	if err != nil {
		return nil, err
	}
	err = message.IsValid()
	if err != nil {
		return nil, err
	}
	err = s.persistence.Save(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (s *MessageService) Get(id string) (*Message, error) {
	return s.persistence.Get(id)
}

func (s *MessageService) Delete(id string) error {
	return s.persistence.Delete(id)
}

func (s *MessageService) GetNext() (*Message, error) {
	return s.persistence.GetNext()
}
