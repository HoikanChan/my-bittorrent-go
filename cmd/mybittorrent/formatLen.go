package main

import (
	"fmt"
	"strconv"
)

const invalidLen = -1

// bit torrent proto need extra 2 char
const protoCharLen = 2

var errGetString = fmt.Errorf("not supported type when getting length")

func getSliceLength(slice []interface{}) (int, error) {
	totalLength := 0
	for _, v := range slice {
		length, err := getStringLength(v)
		if err != nil {
			return invalidLen, err
		}
		totalLength += length
	}
	return totalLength + protoCharLen, nil
}

func getStringLength(str interface{}) (int, error) {
	switch v := str.(type) {
	case string:
		return len(v) + protoCharLen, nil
	case int:
		return len(strconv.Itoa(v)) + protoCharLen, nil
	case []interface{}:
		return getSliceLength(v)
	case map[string]interface{}:
		return getMapLength(v)
	default:
		return invalidLen, errGetString
	}
}

func getMapLength(v map[string]interface{}) (int, error) {
	mapLen := 0
	for key, value := range v {
		valueLen, err := getStringLength(value)
		if err != nil {
			return invalidLen, err
		}
		keyLen, err := getStringLength(key)
		if err != nil {
			return invalidLen, err
		}
		mapLen = mapLen + valueLen + keyLen
	}
	return mapLen + protoCharLen, nil
}
