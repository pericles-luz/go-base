package utils_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestMapInterfaceToByteMustRunCorretly(t *testing.T) {
	// Given
	data := map[string]interface{}{
		"cpf":       utils.TEST_CPF,
		"meioEnvio": utils.SEND_MEDIA_SMS,
	}
	// When
	bytes := utils.MapInterfaceToBytes(data)
	// Then
	require.NotEmpty(t, bytes)
}

func TestMapInterfaceToByteMustReturnNilIfDataIsNil(t *testing.T) {
	// Given
	var data map[string]interface{}
	// When
	bytes := utils.MapInterfaceToBytes(data)
	// Then
	require.Nil(t, bytes)
}

func TestByteToMapInterfaceMustRunCorretly(t *testing.T) {
	// Given
	data := map[string]interface{}{
		"cpf":       utils.TEST_CPF,
		"meioEnvio": utils.SEND_MEDIA_SMS,
	}
	bytes := utils.MapInterfaceToBytes(data)
	// When
	data2 := utils.ByteToMapInterface(bytes)
	// Then
	require.NotEmpty(t, data2)
}

func TestByteToMapInterfaceMustReturnNilIfBytesIsNil(t *testing.T) {
	// Given
	var bytes []byte
	// When
	data := utils.ByteToMapInterface(bytes)
	// Then
	require.Nil(t, data)
}

func TestByteToMapInterfaceMustReturnNilIfBytesIsEmpty(t *testing.T) {
	// Given
	bytes := []byte{}
	// When
	data := utils.ByteToMapInterface(bytes)
	// Then
	require.Nil(t, data)
}

func TestByteToMapInterfaceMustReturnNilIfBytesIsNotJson(t *testing.T) {
	// Given
	bytes := []byte("not json")
	// When
	data := utils.ByteToMapInterface(bytes)
	// Then
	require.Nil(t, data)
}
