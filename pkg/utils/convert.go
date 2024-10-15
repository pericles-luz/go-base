package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Convert a map[string]interface{} to a byte sequence
func MapInterfaceToBytes(data map[string]interface{}) []byte {
	if data == nil {
		return nil
	}
	bytes, _ := json.Marshal(data)
	return bytes
}

// Convert a byte sequence to a map[string]string
func ByteToMapInterface(bytes []byte) map[string]interface{} {
	var data map[string]interface{}
	json.Unmarshal(bytes, &data)
	return data
}

// Convert a map[string]string to a map[string]interface{}
func MapStringToMapInterface(data map[string]string) map[string]interface{} {
	if data == nil {
		return nil
	}
	result := make(map[string]interface{})
	for key, value := range data {
		result[key] = value
	}
	return result
}

// Convert a struct to a map[string]interface{}
// If some error occurs, return nil and the error
func StructToMapInterface(data interface{}) (map[string]interface{}, error) {
	inter, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(inter, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Convert an interface to an int
func InterfaceToInt(incomming interface{}) int {
	if incomming == nil {
		return 0
	}
	switch in := incomming.(type) {
	case int:
		return incomming.(int)
	case int8:
		return int(incomming.(int8))
	case int16:
		return int(incomming.(int16))
	case int32:
		return int(incomming.(int32))
	case int64:
		return int(incomming.(int64))
	case uint:
		return int(incomming.(uint))
	case uint8:
		return int(incomming.(uint8))
	case uint16:
		return int(incomming.(uint16))
	case uint32:
		return int(incomming.(uint32))
	case uint64:
		return int(incomming.(uint64))
	case float32:
		return int(incomming.(float32))
	case float64:
		return int(incomming.(float64))
	case string:
		// if has decimal ponit remove decimal part
		if strings.Contains(incomming.(string), ".") {
			incomming = strings.Split(incomming.(string), ".")[0]
		}
		return StringToInt(incomming.(string))
	default:
		log.Println("InterfaceToInt: unknown type", in)
	}
	return 0
}

// Convert a string to an int
func StringToInt(in string) int {
	var result int
	json.Unmarshal([]byte(strings.TrimLeft(in, "0")), &result)
	return result
}

// Convert an int to a string
func IntToString(in int) string {
	return fmt.Sprintf("%d", in)
}

// Convert a byte sequence to a struct
// If some error occurs, return the error
func ByteToStruct(raw []byte, result interface{}) error {
	return json.Unmarshal(raw, result)
}

// Convert a whatsapp number to a brazilian cellphonenumber
func WhatsappNumberToBrazilianPhonenumber(in string) string {
	if len(in) != WHATSAPP_PHONENUMBER_LENGTH {
		return ""
	}
	ddd := in[2:4]
	phone := in[4:]
	return ddd + "9" + phone
}

// Convert an integer to currency as 1.000.000,00
func IntToCurrency(value uint) string {
	decimal := value % 100
	integer := (value - (value % 100)) / 100
	if integer < 1000 {
		return fmt.Sprintf("%d,%02d", integer, decimal)
	}
	if integer < 1000000 {
		return fmt.Sprintf("%d.%03d,%02d", integer/1000, integer%1000, decimal)
	}
	return fmt.Sprintf("%d.%03d.%03d,%02d", integer/1000000, (integer%1000000)/1000, integer%1000, decimal)
}

// Convert an integer to extense text
func IntToExtense(value int) string {
	extense := []string{
		"zero", "um", "dois", "tres", "quatro", "cinco", "seis", "sete", "oito", "nove",
		"dez", "onze", "doze", "treze", "quatorze", "quinze", "dezesseis", "dezessete", "dezoito", "dezenove",
		"vinte", "trinta", "quarenta", "cinquenta", "sessenta", "setenta", "oitenta", "noventa",
		"cento", "duzentos", "trezentos", "quatrocentos", "quinhentos", "seiscentos", "setecentos", "oitocentos", "novecentos",
	}
	if value < 0 {
		return "menos " + IntToExtense(-value)
	}
	if value == 0 {
		return "zero"
	}
	if value < 20 {
		return extense[value]
	}
	if value < 100 {
		if value%10 == 0 {
			return extense[20+value/10-2]
		}
		return extense[20+value/10-2] + " e " + extense[value%10]
	}
	if value == 100 {
		return "cem"
	}
	if value < 200 {
		return "cento e " + IntToExtense(value%100)
	}
	if value < 1000 {
		if value%100 == 0 {
			return extense[30+value/100-3]
		}
		return extense[30+value/100-3] + " e " + IntToExtense(value%100)
	}
	if value == 1000 {
		return "mil"
	}
	if value < 2000 {
		return "mil e " + IntToExtense(value%1000)
	}
	if value < 1000000 {
		if value%1000 == 0 {
			return IntToExtense(value/1000) + " mil"
		}
		return IntToExtense(value/1000) + " mil e " + IntToExtense(value%1000)
	}
	if value == 1000000 {
		return "um milhão"
	}
	if value < 2000000 {
		return "um milhão e " + IntToExtense(value%1000000)
	}
	if value < 1000000000 {
		if value%1000000 == 0 {
			return IntToExtense(value/1000000) + " milhões"
		}
		return IntToExtense(value/1000000) + " milhões e " + IntToExtense(value%1000000)
	}
	return "número muito grande"
}

// Convert an integer to extense text with cents
func IntToExtenseWithCents(value int) string {
	if value == 0 {
		return "zero reais"
	}
	centsText := "centavos"
	currencyText := "reais"
	integer := value / 100
	cents := value % 100
	if cents == 1 {
		centsText = "centavo"
	}
	if integer == 1 {
		currencyText = "real"
	}
	if integer == 0 {
		return IntToExtense(cents) + " " + centsText
	}
	if cents == 0 {
		return IntToExtense(integer) + " " + currencyText
	}
	return IntToExtense(integer) + " " + currencyText + " e " + IntToExtense(cents) + " " + centsText
}
