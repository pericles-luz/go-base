package utils

import (
	"encoding/json"
	"fmt"
	"log"
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
		return StringToInt(incomming.(string))
	default:
		log.Println("InterfaceToInt: unknown type", in)
	}
	return 0
}

// Convert a string to an int
func StringToInt(in string) int {
	var result int
	json.Unmarshal([]byte(in), &result)
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
