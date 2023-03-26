package utils

import "encoding/json"

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

func InterfaceToInt(in interface{}) int {
	switch in.(type) {
	case int:
		return in.(int)
	case int8:
		return int(in.(int8))
	case int16:
		return int(in.(int16))
	case int32:
		return int(in.(int32))
	case int64:
		return int(in.(int64))
	case uint:
		return int(in.(uint))
	case uint8:
		return int(in.(uint8))
	case uint16:
		return int(in.(uint16))
	case uint32:
		return int(in.(uint32))
	case uint64:
		return int(in.(uint64))
	case float32:
		return int(in.(float32))
	case float64:
		return int(in.(float64))
	case string:
		return StringToInt(in.(string))
	}
	return 0
}

func StringToInt(in string) int {
	var result int
	json.Unmarshal([]byte(in), &result)
	return result
}
