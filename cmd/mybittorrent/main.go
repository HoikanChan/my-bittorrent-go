package main

import (
	// Uncomment this line to pass the first stage
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"unicode"
	// bencode "github.com/jackpal/bencode-go" // Available if you need it!
)

const E = 'e'

// Example:
// - 5:hello -> hello
// - 10:hello12345 -> hello12345
var errInvalidFormat = fmt.Errorf("invalid format")
var errInvalidPrefix = fmt.Errorf("invalid format prefix")
var errInvalidDictKey = fmt.Errorf("invalid dict, key must be string")

func decodeBencode(bencodedString string) (interface{}, error) {
	firstLetter := bencodedString[0]
	switch firstLetter {
	case 'i':
		return decodeToNum(bencodedString)
	case 'l':
		return decodeToList(bencodedString)
	case 'd':
		return decodeToMap(bencodedString)
	default:
		if unicode.IsDigit(rune(firstLetter)) {
			return decodeToStr(bencodedString)
		} else {
			return "", errInvalidPrefix
		}
	}
}

func decodeToMap(bencodedString string) (interface{}, error) {
	lastIdx := len(bencodedString) - 1
	if bencodedString[lastIdx] != E {
		return "", errInvalidFormat
	}
	listContent := bencodedString[1 : len(bencodedString)-1]
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
	if err := iterateParts(listContent, iterator); err != nil {
		return nil, err
	}
	return target, nil
}

func decodeToList(bencodedString string) (interface{}, error) {
	lastIdx := len(bencodedString) - 1
	if bencodedString[lastIdx] != E {
		return "", errInvalidFormat
	}
	listContent := bencodedString[1 : len(bencodedString)-1]

	var target []interface{}
	iterator := func(part interface{}) error {
		target = append(target, part)
		return nil
	}
	if err := iterateParts(listContent, iterator); err != nil {
		return nil, err
	}
	return target, nil
}

func iterateParts(listContent string, iterator func(interface{}) error) error {
	for len(listContent) != 0 {
		d, e := decodeBencode(listContent)
		if e != nil {
			return e
		}
		strLen, e := getStringLength(d)
		if e != nil {
			return e
		}
		if e := iterator(d); e != nil {
			return e
		}
		listContent = listContent[strLen:]
	}
	return nil
}

func decodeToNum(bencodedString string) (interface{}, error) {
	firstEndIndex, err := findChar(bencodedString, E)
	if err != nil {
		return "", errInvalidFormat
	}

	i, err := strconv.Atoi(bencodedString[1:firstEndIndex])
	if err != nil {
		return "", err
	}

	return i, nil
}

func decodeToStr(bencodedString string) (interface{}, error) {
	firstColonIndex, err := findChar(bencodedString, ':')
	if err != nil {
		return "", errInvalidFormat
	}

	lengthStr := bencodedString[:firstColonIndex]

	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return "", err
	}

	return bencodedString[firstColonIndex+1 : firstColonIndex+1+length], nil
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

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	command := os.Args[1]

	if command == "decode" {
		// Uncomment this block to pass the first stage
		//
		bencodedValue := os.Args[2]

		decoded, err := decodeBencode(bencodedValue)
		if err != nil {
			fmt.Println(err)
			return
		}

		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))
	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
