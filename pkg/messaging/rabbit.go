package messaging

import (
	"context"
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/pericles-luz/go-base/internals/migration"
	"github.com/pericles-luz/go-base/pkg/conf"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	reconnectDelay    = 5 * time.Second
	resendDelay       = 5 * time.Second
	MESSAGING_TIMEOUT = 30
)

type Rabbit struct {
	conn          *amqp.Connection
	ch            *amqp.Channel
	done          chan bool
	notifyClose   chan *amqp.Error
	notifyConfirm chan amqp.Confirmation
	isConnected   bool
	dsn           string
}

type RabbitConfig struct {
	file conf.ConfigBase
	DSN  string `json:"dsn"`
}

func NewRabbit(file string) *Rabbit {
	config := RabbitConfig{}
	err := config.LoadRabbitConfig(file)
	if err != nil {
		panic(err)
	}
	rabbit := Rabbit{
		dsn:  config.DSN,
		done: make(chan bool),
	}
	go rabbit.handleReconnect()
	return &rabbit
}

func (r *Rabbit) Publish(exchange, routingKey string, body []byte) error {
	if !r.isConnected {
		return errors.New("Rabbit is not connected")
	}
	ctx := context.Background()
	err := r.ch.PublishWithContext(ctx, exchange, routingKey, false, false, amqp.Publishing{Body: body})
	if err != nil {
		return err
	}
	return nil
}

func (r *Rabbit) Consume(queue string, callback func(amqp.Delivery)) error {
	maxTry := 3
	for !r.isConnected {
		log.Println("RabbitMQ: not connected")
		time.Sleep(reconnectDelay)
		maxTry--
	}
	var incoming chan struct{}
	msgs, err := r.ch.Consume(queue, "", false, false, false, false, nil)
	if err != nil {
		return err
	}
	go func() {
		for d := range msgs {
			callback(d)
		}
	}()
	<-incoming
	return nil
}

func (r *Rabbit) DeclareExchange(exchange string) error {
	if !r.isConnected {
		return errors.New("Rabbit is not connected")
	}
	log.Println("RabbitMQ: declaring exchange", exchange)
	err := r.ch.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	if err != nil {
		return err
	}
	log.Println("RabbitMQ: declared exchange", exchange)
	return nil
}

func (r *Rabbit) DeclareQueue(queue, exchange, routingKey string) error {
	if !r.isConnected {
		return errors.New("Rabbit is not connected")
	}
	log.Println("RabbitMQ: declaring queue", queue)
	_, err := r.ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}
	log.Println("RabbitMQ: declared queue", queue)
	log.Println("RabbitMQ: binding queue", queue, "to exchange", exchange, "with routing key", routingKey)
	err = r.ch.QueueBind(queue, routingKey, exchange, false, nil)
	if err != nil {
		return err
	}
	log.Println("RabbitMQ: bound queue", queue, "to exchange", exchange, "with routing key", routingKey)
	return nil
}

func (r *RabbitConfig) LoadRabbitConfig(file string) error {
	data, err := r.file.ReadConfigurationFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, r)
	if err != nil {
		return err
	}
	return nil
}

func (r *Rabbit) Disconnect() error {
	err := r.ch.Close()
	if err != nil {
		return err
	}
	err = r.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func (r *Rabbit) handleReconnect() {
	for {
		r.isConnected = false
		log.Println("RabbitMQ: connecting")
		for !r.connect() {
			log.Println("RabbitMQ: reconnecting in", reconnectDelay)
			time.Sleep(reconnectDelay)
		}
		select {
		case <-r.done:
			return
		case <-r.notifyClose:
			log.Println("RabbitMQ: closed")
		}
	}
}

func (r *Rabbit) connect() bool {
	conn, err := amqp.DialConfig(r.dsn, amqp.Config{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})
	// conn, err := amqp.Dial(r.dsn)
	if err != nil {
		log.Println("RabbitMQ: dialing:", err)
		return false
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Println("RabbitMQ: channel:", err)
		return false
	}
	r.conn = conn
	r.ch = ch
	r.isConnected = true
	r.notifyClose = make(chan *amqp.Error)
	r.notifyConfirm = make(chan amqp.Confirmation)
	r.ch.NotifyClose(r.notifyClose)
	r.ch.NotifyPublish(r.notifyConfirm)
	log.Println("RabbitMQ: connected")
	return true
}

func (r *Rabbit) IsConnected() bool {
	return r.isConnected
}

func (r *Rabbit) PublishFromCache(messageService *migration.MessageService) error {
	for {
		message, err := messageService.GetNext()
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if message == nil {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		err = r.Publish(message.GetExchange(), message.GetRoutingKey(), []byte(message.GetData()))
		if err != nil {
			return err
		}
		err = messageService.Delete(message.GetID())
		if err != nil {
			return err
		}
	}
}
