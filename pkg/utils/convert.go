package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

func MapInterfaceToBytes(data map[string]interface{}) []byte {
	if data == nil {
		return nil
	}
	bytes, _ := json.Marshal(data)
	return bytes
}

func ByteToMapInterface(bytes []byte) map[string]interface{} {
	var data map[string]interface{}
	json.Unmarshal(bytes, &data)
	return data
}

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

func StringToInt(in string) int {
	var result int
	json.Unmarshal([]byte(in), &result)
	return result
}

func IntToString(in int) string {
	return fmt.Sprintf("%d", in)
}

func ByteToStruct(raw []byte, result interface{}) error {
	return json.Unmarshal(raw, result)
}

func WhatsappNumberToBrazilianPhonenumber(in string) string {
	if len(in) != WHATSAPP_PHONENUMBER_LENGTH {
		return ""
	}
	ddd := in[2:4]
	phone := in[4:]
	return ddd + "9" + phone
}
