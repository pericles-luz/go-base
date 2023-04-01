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

func TestStringToIntMustConvertCorrectly(t *testing.T) {
	// Given
	value := "123"
	// When
	result := utils.StringToInt(value)
	// Then
	require.Equal(t, 123, result)
}

func TestStringToIntMustReturnZeroIfValueIsNotNumber(t *testing.T) {
	// Given
	value := "not number"
	// When
	result := utils.StringToInt(value)
	// Then
	require.Equal(t, 0, result)
}

func TestInterfaceToIntMustConvertStringToInt(t *testing.T) {
	// Given
	value := "123"
	// When
	result := utils.InterfaceToInt(value)
	// Then
	require.Equal(t, 123, result)
}

func TestInterfaceToIntMustConvertIntToInt(t *testing.T) {
	// Given
	value := 123
	// When
	result := utils.InterfaceToInt(value)
	// Then
	require.Equal(t, 123, result)
}

func TestWhatsappNumberToBrazilianPhonenumber(t *testing.T) {
	// Given
	number := "553186058910"
	// When
	result := utils.WhatsappNumberToBrazilianPhonenumber(number)
	// Then
	require.Equal(t, "31986058910", result)
}
