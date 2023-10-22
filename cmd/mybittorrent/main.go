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
var errFormat = fmt.Errorf("invalid format")

func decodeBencode(bencodedString string) (interface{}, error) {
	firstLetter := bencodedString[0]
	if unicode.IsDigit(rune(firstLetter)) {
		firstColonIndex, err := findChar(bencodedString, ':')
		if err != nil {
			return "", errFormat
		}

		lengthStr := bencodedString[:firstColonIndex]

		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return "", err
		}

		return bencodedString[firstColonIndex+1 : firstColonIndex+1+length], nil
	} else if firstLetter == 'i' {

		firstEndIndex, err := findChar(bencodedString, E)
		if err != nil {
			return "", errFormat
		}

		i, err := strconv.Atoi(bencodedString[1:firstEndIndex])
		if err != nil {
			return "", err
		}

		return i, nil
	} else if firstLetter == 'l' {
		var target []interface{}
		listContent := bencodedString[1 : len(bencodedString)-1]
		for len(listContent) != 0 {
			d, e := decodeBencode(listContent)
			if e != nil {
				return "", e
			}
			strLen, e := getStringLength(d)
			if e != nil {
				return "", e
			}
			target = append(target, d)
			listContent = listContent[strLen:]
		}
		return target, nil
	} else {
		return "", fmt.Errorf("only strings are supported at the moment")
	}
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
