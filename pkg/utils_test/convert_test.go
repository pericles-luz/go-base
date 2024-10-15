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
		"sentMedia": utils.SEND_MEDIA_SMS,
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
		"sentMedia": utils.SEND_MEDIA_SMS,
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

func TestStringToIntMustConvertCorrectlyWhenStringStartsWithZero(t *testing.T) {
	// Given
	value := "00000123"
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

func TestInterfaceToIntMustConvertFloatStringToInt(t *testing.T) {
	// Given
	value := "123.45"
	// When
	result := utils.InterfaceToInt(value)
	// Then
	require.Equal(t, 123, result)
	// Given
	value = "123.0"
	// When
	result = utils.InterfaceToInt(value)
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

func TestIntToString(t *testing.T) {
	// Given
	value := 123
	// When
	result := utils.IntToString(value)
	// Then
	require.Equal(t, "123", result)
}

func TestIntToCurrency(t *testing.T) {
	// Given
	value := 1234567
	// When
	result := utils.IntToCurrency(uint(value))
	// Then
	require.Equal(t, "12.345,67", result)
}

func TestIntToCurrencyMustReturnZeroStringIfValueIsZero(t *testing.T) {
	// Given
	value := 0
	// When
	result := utils.IntToCurrency(uint(value))
	// Then
	require.Equal(t, "0,00", result)
}

func TestIntToCurrencyMustReturnCorrectWhenValueIsOneCent(t *testing.T) {
	// Given
	value := 1
	// When
	result := utils.IntToCurrency(uint(value))
	// Then
	require.Equal(t, "0,01", result)
}

func TestIntToExtenseMustReturnCorrectlyWhenValueIsLessThanTwenty(t *testing.T) {
	require.Equal(t, "zero", utils.IntToExtense(0))
	require.Equal(t, "um", utils.IntToExtense(1))
	require.Equal(t, "dez", utils.IntToExtense(10))
	require.Equal(t, "quinze", utils.IntToExtense(15))
	require.Equal(t, "dezenove", utils.IntToExtense(19))
}

func TestIntToExtenseMustReturnCorrectlyWhenValueIsLessThanHundred(t *testing.T) {
	require.Equal(t, "vinte", utils.IntToExtense(20))
	require.Equal(t, "vinte e um", utils.IntToExtense(21))
	require.Equal(t, "trinta e cinco", utils.IntToExtense(35))
	require.Equal(t, "quarenta e nove", utils.IntToExtense(49))
	require.Equal(t, "cinquenta", utils.IntToExtense(50))
	require.Equal(t, "sessenta e tres", utils.IntToExtense(63))
	require.Equal(t, "setenta e sete", utils.IntToExtense(77))
	require.Equal(t, "oitenta e nove", utils.IntToExtense(89))
}

func TestIntToExtenseMustReturnCorrectlyWhenValueIsLessThanThousandAndGreaterThanHundred(t *testing.T) {
	require.Equal(t, "cem", utils.IntToExtense(100))
	require.Equal(t, "cento e um", utils.IntToExtense(101))
	require.Equal(t, "cento e trinta e cinco", utils.IntToExtense(135))
	require.Equal(t, "duzentos", utils.IntToExtense(200))
	require.Equal(t, "duzentos e um", utils.IntToExtense(201))
	require.Equal(t, "duzentos e quinze", utils.IntToExtense(215))
	require.Equal(t, "quinhentos", utils.IntToExtense(500))
	require.Equal(t, "quinhentos e trinta e um", utils.IntToExtense(531))
	require.Equal(t, "novecentos", utils.IntToExtense(900))
	require.Equal(t, "novecentos e setenta e um", utils.IntToExtense(971))
}

func TestIntToExtenseMustReturnCorrectlyWhenValueIsLessThanMillion(t *testing.T) {
	require.Equal(t, "mil", utils.IntToExtense(1000))
	require.Equal(t, "mil e um", utils.IntToExtense(1001))
	require.Equal(t, "mil e trinta e cinco", utils.IntToExtense(1035))
	require.Equal(t, "dois mil", utils.IntToExtense(2000))
	require.Equal(t, "dois mil e um", utils.IntToExtense(2001))
	require.Equal(t, "dois mil e quinze", utils.IntToExtense(2015))
	require.Equal(t, "quinhentos mil", utils.IntToExtense(500000))
	require.Equal(t, "quinhentos mil e trinta e um", utils.IntToExtense(500031))
	require.Equal(t, "novecentos e noventa e nove mil", utils.IntToExtense(999000))
	require.Equal(t, "novecentos e noventa e nove mil e novecentos e noventa e nove", utils.IntToExtense(999999))
}

func TestIntToExtenseMustReturnCorrectlyWhenValueIsLessThanBillion(t *testing.T) {
	require.Equal(t, "um milhão", utils.IntToExtense(1000000))
	require.Equal(t, "um milhão e um", utils.IntToExtense(1000001))
	require.Equal(t, "um milhão e trinta e cinco", utils.IntToExtense(1000035))
	require.Equal(t, "dois milhões", utils.IntToExtense(2000000))
	require.Equal(t, "dois milhões e um", utils.IntToExtense(2000001))
	require.Equal(t, "dois milhões e quinze", utils.IntToExtense(2000015))
	require.Equal(t, "quinhentos milhões", utils.IntToExtense(500000000))
	require.Equal(t, "quinhentos milhões e trinta e um", utils.IntToExtense(500000031))
	require.Equal(t, "novecentos e noventa e nove milhões", utils.IntToExtense(999000000))
	require.Equal(t, "novecentos e noventa e nove milhões e novecentos e noventa e nove mil e novecentos e noventa e nove", utils.IntToExtense(999999999))
}

func TestIntToExtenseMustReturnCorrectlyWhenValueIsGreaterThanBillion(t *testing.T) {
	require.Equal(t, "número muito grande", utils.IntToExtense(1000000000))
}

func TestIntToExtenseWithCentsMustReturnCorrectlyWhenLessThanOne(t *testing.T) {
	require.Equal(t, "zero reais", utils.IntToExtenseWithCents(0))
	require.Equal(t, "um centavo", utils.IntToExtenseWithCents(1))
	require.Equal(t, "quinze centavos", utils.IntToExtenseWithCents(15))
	require.Equal(t, "dezenove centavos", utils.IntToExtenseWithCents(19))
	require.Equal(t, "vinte centavos", utils.IntToExtenseWithCents(20))
	require.Equal(t, "vinte e um centavos", utils.IntToExtenseWithCents(21))
	require.Equal(t, "trinta e cinco centavos", utils.IntToExtenseWithCents(35))
}

func TestIntToExtenseWithCentsMustReturnCorrectlyWhenLessThanOneThousand(t *testing.T) {
	require.Equal(t, "um real", utils.IntToExtenseWithCents(100))
	require.Equal(t, "um real e trinta e cinco centavos", utils.IntToExtenseWithCents(135))
	require.Equal(t, "cem reais", utils.IntToExtenseWithCents(10000))
	require.Equal(t, "cento e um reais", utils.IntToExtenseWithCents(10100))
	require.Equal(t, "cento e trinta e cinco reais", utils.IntToExtenseWithCents(13500))
	require.Equal(t, "duzentos reais", utils.IntToExtenseWithCents(20000))
	require.Equal(t, "duzentos e um reais", utils.IntToExtenseWithCents(20100))
	require.Equal(t, "duzentos e quinze reais", utils.IntToExtenseWithCents(21500))
	require.Equal(t, "quinhentos reais", utils.IntToExtenseWithCents(50000))
	require.Equal(t, "quinhentos e trinta e um reais", utils.IntToExtenseWithCents(53100))
	require.Equal(t, "novecentos reais", utils.IntToExtenseWithCents(90000))
	require.Equal(t, "novecentos e setenta e um reais", utils.IntToExtenseWithCents(97100))
}
