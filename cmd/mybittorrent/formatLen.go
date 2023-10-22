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
	return totalLength, nil
}

func getStringLength(str interface{}) (int, error) {
	switch v := str.(type) {
	case string:
		return len(v) + protoCharLen, nil
	case int:
		return len(strconv.Itoa(v)) + protoCharLen, nil
	case []interface{}:
		sliceLen, err := getSliceLength(v)
		if err != nil {
			return invalidLen, err
		}
		return sliceLen + protoCharLen, nil
	default:
		return invalidLen, errGetString
	}
}
