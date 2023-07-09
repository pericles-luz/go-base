package messaging_test

import (
	"context"
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
		log.Println("Message received", d.DeliveryTag)
	}
	rabbit.Consume("ct.teste", proccess)
	time.Sleep(3 * time.Second)
}

func TestRabbitPublishConsumeWithContext(t *testing.T) {
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
		log.Println("Message received", d.DeliveryTag)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	rabbit.ConsumeWithContext(ctx, "ct.teste", proccess)
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
	for i := 0; i < 5; i++ {
		messageService.Send("tst.teste", "teste", fmt.Sprintf(`{"teste":"testado de cache %v"}`, i), 1)
	}
	rabbit := messaging.NewRabbit("rabbit")
	time.Sleep(time.Second * 3)
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
