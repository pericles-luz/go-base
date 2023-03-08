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
