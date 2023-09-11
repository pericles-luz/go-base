package messaging_test

import (
	"testing"

	"github.com/pericles-luz/go-base/internal/migration"
	"github.com/stretchr/testify/require"
)

// mockgen -destination=application/mocks/message.go -source=adapter/messaging/message.go application

func TestMessage_IsValid(t *testing.T) {
	message := migration.NewMessage()
	// should validate a message with all fields informed
	message.SetExchange("testExchange")
	message.SetRoutingKey("testRoutingKey")
	message.SetData("testData")
	message.SetDurable(1)
	err := message.IsValid()
	require.NoError(t, err)
	// should not validate a message with exchange not informed
	message.SetExchange("")
	err = message.IsValid()
	require.EqualError(t, err, migration.EXCHANGE_NOT_INFORMED)
	// should not validate a message with routing key not informed
	message.SetExchange("testExchange")
	message.SetRoutingKey("")
	err = message.IsValid()
	require.EqualError(t, err, migration.ROUTING_KEY_NOT_INFORMED)
	// should not validate a message with data not informed
	message.SetExchange("testExchange")
	message.SetRoutingKey("testRoutingKey")
	message.SetData("")
	err = message.IsValid()
	require.EqualError(t, err, migration.DATA_NOT_INFORMED)
}
