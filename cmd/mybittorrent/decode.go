package main

import (
	"fmt"
	"strconv"
	"unicode"
)

const E = 'e'
var errInvalidFormat = fmt.Errorf("invalid format")
var errInvalidPrefix = fmt.Errorf("invalid format prefix")
var errInvalidDictKey = fmt.Errorf("invalid dict, key must be string")
var errEndOfFile = fmt.Errorf("file ended")
var errNoEnd = fmt.Errorf("can not find end of list or map")
var errInvalidNum = fmt.Errorf("invalid number")

// Example:
// - 5:hello -> hello
// - 10:hello12345 -> hello12345
func decodeBencode(bencodedString string) (interface{}, int, error) {
	firstLetter := bencodedString[0]
	switch firstLetter {
	case 'i':
		return decodeToNum(bencodedString)
	case 'l':
		return decodeToList(bencodedString)
	case 'd':
		return decodeToMap(bencodedString)
	case 'e':
		return "", 0, errEndOfFile
	default:
		if unicode.IsDigit(rune(firstLetter)) {
			return decodeToStr(bencodedString)
		} else {
			return "", 0, errInvalidPrefix
		}
	}
}

func decodeToMap(bencodedString string) (interface{}, int, error) {
	var key string
	target := make(map[string]interface{})
	iterator := func(part interface{}) error {
		if key == "" {
			if s, success := part.(string); success {
				key = s
			} else {
				return errInvalidDictKey
			}
		} else {
			target[key] = part
			// reset key after setting value
			key = ""
		}
		return nil
	}
	iteratedLen, err := iterate(bencodedString, iterator)
	if err != nil {
		return nil, 0, err
	}
	return target, iteratedLen, nil
}

func decodeToList(bencodedString string) (interface{}, int, error) {
	var target []interface{}
	iterator := func(part interface{}) error {
		target = append(target, part)
		return nil
	}
	iteratedLen, err := iterate(bencodedString, iterator)
	if err != nil {
		return nil, 0, err
	}
	return target, iteratedLen, nil
}

// iterate the map or list of bittorrent
func iterate(bencodedString string, iterator func(interface{}) error) (int, error) {
	content := bencodedString[1:]
	iteratedLen := 0
	for len(content) != 0 {
		d, strLen, e := decodeBencode(content)
		if e != nil {
			if e == errEndOfFile {
				return iteratedLen + 1, nil
			}
			return 0, e
		}
		if e := iterator(d); e != nil {
			return 0, e
		}
		content = content[strLen:]
		iteratedLen += strLen
	}
	return 0, errNoEnd
}

func decodeToNum(bencodedString string) (interface{}, int, error) {
	firstEndIndex, err := findChar(bencodedString, E)
	if err != nil {
		return "", 0, errInvalidFormat
	}

	i, err := strconv.Atoi(bencodedString[1:firstEndIndex])
	if err != nil {
		return "", 0, errInvalidNum
	}

	return i, firstEndIndex + 1, nil
}

func decodeToStr(bencodedString string) (interface{}, int, error) {
	firstColonIndex, err := findChar(bencodedString, ':')
	if err != nil {
		return "", 0, errInvalidFormat
	}

	lengthStr := bencodedString[:firstColonIndex]

	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return "", 0, err
	}

	return bencodedString[firstColonIndex+1 : firstColonIndex+1+length], firstColonIndex + 1 + length, nil
}

func findChar(bencodedString string, target byte) (result int, err error) {
	result = -1
	for i := 0; i < len(bencodedString); i++ {
		if bencodedString[i] == target {
			result = i
			break
		}
	}
	if result == -1 {
		err = fmt.Errorf("can't find target char")
	}
	return result, err
}
