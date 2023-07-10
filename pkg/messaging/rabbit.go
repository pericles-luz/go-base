package messaging

import (
	"context"
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/pericles-luz/go-base/internals/factory"
	"github.com/pericles-luz/go-base/internals/migration"
	"github.com/pericles-luz/go-base/pkg/conf"
	"github.com/pericles-luz/go-base/pkg/infra/database"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	reconnectDelay    = 5 * time.Second
	resendDelay       = 5 * time.Second
	MESSAGING_TIMEOUT = 30
	MAX_RETRIES       = 3
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

// NewRabbitPublisher is a function that creates a new RabbitMQ publisher
// and publishes all messages from the database.
// It is used to publish messages that were not published due to a RabbitMQ
// connection failure.
// The messages are published unordered
// The table used to store the messages is:
// create table if not exists RabbitCache(RabbitCacheID string primary key, DE_Exchange string, DE_RoutingKey string, JS_Data text, SN_Durable integer, TS_Operacao string, ID_Status integer default 0)
// The table is created in memory and is not persisted.
// To send a message, you have to call the function Send of the message service that uses the same database connection.
// the function needs the following parameters:
// - exchange: the exchange name
// - routingKey: the routing key
// - data: the message data
// - durable: if the message is durable
// CAUTION: if you use more than one dispatcher, the messages can be sent more than once. You have to take care of this.
// The function returns a message service that can be used to send messages.
func NewRabbitPublisher(file string, pool *database.Pool, mtx *sync.Mutex, dispatchers int) *migration.MessageService {
	messageService := factory.NewMessageService(pool)

	for i := 0; i < dispatchers; i++ {
		rabbit := NewRabbit(file)
		count := 5
		for count > 0 && !rabbit.IsConnected() {
			log.Println("RabbitMQ: trying to connect...")
			time.Sleep(1 * time.Second)
			count--
		}
		if !rabbit.IsConnected() {
			log.Println("RabbitMQ: not connected")
			return messageService
		}
		go rabbit.PublishFromCache(messageService, mtx)
	}
	return messageService
}

func (r *Rabbit) Publish(exchange, routingKey string, body []byte) error {
	if !r.isConnected {
		return errors.New("Rabbit is not connected")
	}
	ctx, cancel := context.WithTimeout(context.Background(), resendDelay)
	defer cancel()
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

func (r *Rabbit) ConsumeWithContext(ctx context.Context, queue string, callback func(amqp.Delivery)) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			err := r.consumeWithContext(ctx, queue, callback)
			if err != nil {
				log.Println("RabbitMQ: error consuming", err)
				time.Sleep(reconnectDelay)
			}
			if err := recover(); err != nil {
				log.Println("RabbitMQ: error consuming", err)
				time.Sleep(reconnectDelay)
			}
		}
	}
}

func (r *Rabbit) ConsumeResilient(queue string, callback func(amqp.Delivery)) error {
	for {
		r.Consume(queue, callback)
		if err := recover(); err != nil {
			log.Println("RabbitMQ: FAILED consuming", err)
			time.Sleep(reconnectDelay)
		}
	}
}

func (r *Rabbit) consumeWithContext(ctx context.Context, queue string, callback func(amqp.Delivery)) error {
	maxTry := 3
	for !r.isConnected {
		log.Println("RabbitMQ: not connected")
		time.Sleep(reconnectDelay)
		maxTry--
	}
	msgs, err := r.ch.Consume(queue, "", false, false, false, false, nil)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case d := <-msgs:
				callback(d)
			}
		}
	}()
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
	err = ch.Qos(1, 0, false)
	if err != nil {
		log.Println("RabbitMQ: qos:", err)
		return false
	}
	// err = ch.Confirm(true)
	// if err != nil {
	// 	log.Println("RabbitMQ: confirm:", err)
	// 	return false
	// }
	r.conn = conn
	r.ch = ch
	r.isConnected = true
	r.notifyClose = make(chan *amqp.Error)
	r.notifyConfirm = make(chan amqp.Confirmation, 100)
	r.ch.NotifyClose(r.notifyClose)
	r.ch.NotifyPublish(r.notifyConfirm)
	err = r.ch.Confirm(false)
	if err != nil {
		log.Println("RabbitMQ: confirm:", err)
		return false
	}
	log.Println("RabbitMQ: connected")
	return true
}

func (r *Rabbit) IsConnected() bool {
	return r.isConnected
}

func (r *Rabbit) PublishFromCache(messageService *migration.MessageService, mtx *sync.Mutex) error {
	for {
		if err := r.publishFromCache(messageService, mtx); err != nil {
			log.Println("RabbitMQ: publish from cache FAILED:", err)
			time.Sleep(500 * time.Millisecond)
			continue
		}
		if err := recover(); err != nil {
			log.Println("RabbitMQ: publish from cache FAILED:", err)
			time.Sleep(500 * time.Millisecond)
			continue
		}
	}
}

func (r *Rabbit) publishFromCache(messageService *migration.MessageService, mtx *sync.Mutex) error {
	for {
		if !r.IsConnected() {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		mtx.Lock()
		message, err := messageService.GetNext()
		mtx.Unlock()
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if message == nil {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		err = r.Publish(message.GetExchange(), message.GetRoutingKey(), []byte(message.GetData()))
		if err != nil {
			log.Println("RabbitMQ: failed to publish from cache", message.GetID())
			return err
		}
		tries := MAX_RETRIES
	waitingLoop:
		for tries > 0 {
			select {
			case confirm := <-r.notifyConfirm:
				if confirm.Ack {
					err = messageService.Delete(message.GetID())
					if err != nil {
						log.Println("RabbitMQ: failed to delete cached", message.GetID())
						return err
					}
				} else {
					log.Println("RabbitMQ: failed to confirm", message.GetID())
					time.Sleep(500 * time.Millisecond)
				}
				break waitingLoop
			case <-time.After(5 * time.Second):
				log.Println("RabbitMQ: timeout waiting for confirmation of cached sending", message.GetID())
				time.Sleep(500 * time.Millisecond)
				tries--
			}
		}
	}
}
