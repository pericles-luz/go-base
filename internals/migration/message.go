package migration

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	EXCHANGE_NOT_INFORMED    = "exchange not informed"
	ROUTING_KEY_NOT_INFORMED = "routing key not informed"
	DATA_NOT_INFORMED        = "data not informed"
)

type Message struct {
	RabbitCacheID string `json:"id"`
	DE_Exchange   string `json:"exchange" validate:"required"`
	DE_RoutingKey string `json:"routingKey" validate:"required"`
	JS_Data       string `json:"data" validate:"required"`
	SN_Durable    int16  `json:"durable" validate:"required"`
	TS_Operacao   time.Time
}

func NewMessage() *Message {
	message := &Message{}
	message.SetID(uuid.New().String())
	message.SetDurable(1)
	return message
}

func (m *Message) SetExchange(exchange string) {
	m.DE_Exchange = exchange
}

func (m *Message) SetRoutingKey(routingKey string) {
	m.DE_RoutingKey = routingKey
}

func (m *Message) SetData(data interface{}) {
	if _, ok := data.(string); ok {
		m.JS_Data = data.(string)
		return
	}
	dataString := toJSON(data)
	m.JS_Data = dataString
}

func (m *Message) SetDurable(durable int16) {
	m.SN_Durable = durable
}

func (m *Message) SetID(id string) {
	m.RabbitCacheID = id
}

func (m *Message) IsValid() error {
	if m.DE_Exchange == "" {
		return errors.New(EXCHANGE_NOT_INFORMED)
	}
	if m.DE_RoutingKey == "" {
		return errors.New(ROUTING_KEY_NOT_INFORMED)
	}
	if m.JS_Data == "" {
		return errors.New(DATA_NOT_INFORMED)
	}
	return nil
}

func (m *Message) GetExchange() string {
	return m.DE_Exchange
}

func (m *Message) GetRoutingKey() string {
	return m.DE_RoutingKey
}

func (m *Message) GetData() string {
	return m.JS_Data
}

func (m *Message) GetDurable() int16 {
	return m.SN_Durable
}

func (m *Message) GetID() string {
	return m.RabbitCacheID
}

func toJSON(msg interface{}) string {
	dataObj := struct {
		Data interface{} `json:"data"`
	}{
		Data: msg,
	}
	r, err := json.Marshal(dataObj.Data)
	if err != nil {
		return err.Error()
	}
	return string(r)
}
