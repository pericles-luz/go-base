package utils

import (
	"encoding/json"
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