package messaging_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/pericles-luz/go-base/internals/factory"
	"github.com/pericles-luz/go-base/internals/migration"
	"github.com/pericles-luz/go-base/pkg/messaging"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/require"
)

func TestRabbitConsume(t *testing.T) {
	t.Skip("Test only if necessary")
	rabbit := messaging.NewRabbit("rabbit")
	time.Sleep(3 * time.Second)
	process := func(d amqp.Delivery) {
		var msg map[string]interface{}
		err := json.Unmarshal(d.Body, &msg)
		if err != nil {
			t.Error(err)
		}
		t.Log(msg, msg["teste"], d.DeliveryTag)
		d.Acknowledger.Ack(d.DeliveryTag, false)
	}
	rabbit.Consume("vt.teste", process)
}

func TestRabbitPublish(t *testing.T) {
	t.Skip("Test only if necessary")
	rabbit := messaging.NewRabbit("rabbit")
	time.Sleep(3 * time.Second)
	msg := map[string]interface{}{
		"teste": "teste",
	}
	body, err := json.Marshal(msg)
	if err != nil {
		t.Error(err)
	}
	err = rabbit.Publish("tst.teste", "teste", body)
	if err != nil {
		t.Error(err)
	}
}

func TestRabbitPublishConsume(t *testing.T) {
	t.Skip("Test only if necessary")
	rabbit := messaging.NewRabbit("rabbit")
	time.Sleep(3 * time.Second)
	msg := map[string]interface{}{
		"teste": "teste",
	}
	body, err := json.Marshal(msg)
	if err != nil {
		t.Error(err)
	}
	err = rabbit.Publish("tst.teste", "teste", body)
	if err != nil {
		t.Error(err)
	}
	proccess := func(d amqp.Delivery) {
		err := receiveMessage(d.Body)
		if err != nil {
			t.Error(err)
			d.Acknowledger.Nack(d.DeliveryTag, false, true)
			return
		}
		d.Acknowledger.Ack(d.DeliveryTag, false)
	}
	rabbit.Consume("vt.teste", proccess)
}

func TestRabbitSendToAnotherQueue(t *testing.T) {
	t.Skip("Test only if necessary")
	rabbit := messaging.NewRabbit("rabbit")
	msg := map[string]interface{}{
		"teste": "teste",
	}
	body, err := json.Marshal(msg)
	if err != nil {
		t.Error(err)
	}
	err = rabbit.Publish("tst.teste", "teste", body)
	if err != nil {
		t.Error(err)
	}
	proccess := func(d amqp.Delivery) {
		var msg map[string]interface{}
		err := json.Unmarshal(d.Body, &msg)
		if err != nil {
			t.Error(err)
		}
		t.Log(msg, msg["teste"], d.DeliveryTag)
		body, err := json.Marshal(msg)
		if err != nil {
			t.Error(err)
		}
		rabbit.Publish("tst.teste", "teste2", body)
		d.Acknowledger.Nack(d.DeliveryTag, false, true)
	}
	rabbit.Consume("vt.teste", proccess)
}

func TestRabbitPublishFromRabbitCache(t *testing.T) {
	t.Skip("Test only if necessary")
	tearDown, pool := migration.SetupTest(t)
	defer tearDown(t)
	messageDB := factory.NewMessageDB(pool)
	messageService := migration.NewMessageService(messageDB)
	messageService.Send("tst.teste", "teste", `{"teste":"testado de cache"}`, 1)
	rabbit := messaging.NewRabbit("rabbit")
	time.Sleep(time.Second * 3)
	defer rabbit.Disconnect()
	go func(service *migration.MessageService, rabbit *messaging.Rabbit) {
		var message *migration.Message
		var err error
		count := 0
		for {
			count++
			fmt.Println("count", count)
			if count > 5 {
				break
			}
			message, err = service.GetNext()
			if err != nil {
				fmt.Println("erro: ", err)
			}
			if message == nil {
				fmt.Println("message nil")
				_, err = service.Send("tst.teste", "teste", `{"teste":"testado 1"}`, 1)
				require.NoError(t, err)
				continue
			}
			err = rabbit.Publish(message.GetExchange(), message.GetRoutingKey(), []byte(message.GetData()))
			require.NoError(t, err)
			fmt.Println("apagando: ", message)
			err = service.Delete(message.GetID())
			require.NoError(t, err)
		}
	}(messageService, rabbit)
	t.Log("esperando para finalizar")
	time.Sleep(time.Second * 3)
}

func TestRabbitServicePublishFromRabbitCache(t *testing.T) {
	t.Skip("Test only if necessary")
	tearDown, pool := migration.SetupTest(t)
	defer tearDown(t)
	messageService := factory.NewMessageService(pool)
	messageService.Send("tst.teste", "teste", `{"teste":"testado de cache 0"}`, 1)
	messageService.Send("tst.teste", "teste", `{"teste":"testado de cache 1"}`, 1)
	messageService.Send("tst.teste", "teste", `{"teste":"testado de cache 2"}`, 1)
	messageService.Send("tst.teste", "teste", `{"teste":"testado de cache 3"}`, 1)
	rabbit := messaging.NewRabbit("rabbit")
	defer rabbit.Disconnect()
	go rabbit.PublishFromCache(messageService)
	t.Log("esperando para finalizar")
	time.Sleep(time.Second * 3)
}

func TestRabbitDeclareExchange(t *testing.T) {
	t.Skip("Use only when needed it")
	rabbit := messaging.NewRabbit("rabbit")
	time.Sleep(2 * time.Second)
	require.NoError(t, rabbit.DeclareExchange("tst.teste"))
	require.NoError(t, rabbit.DeclareQueue("ct.teste", "tst.teste", "teste"))
	rabbit.Disconnect()
}

func receiveMessage(body []byte) error {
	var msg map[string]interface{}
	err := json.Unmarshal(body, &msg)
	if err != nil {
		return err
	}
	log.Println(msg, msg["teste"])
	return nil
}
