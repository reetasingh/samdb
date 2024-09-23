package resp

import (
	"fmt"
	"strconv"
)

func decodeInt(input []byte) (int, int, error) {
	var value int
	if len(input) <= 3 {
		return 0, 0, fmt.Errorf("not a valid integer input. length < 3")
	}
	if input[0] != ':' {
		return 0, 0, fmt.Errorf("not a valid integer input. should start with :")
	}
	if input[1] != '+' && input[1] != '-' {
		return 0, 0, fmt.Errorf("not a valid integer input. should start with :")
	}
	start := 2
	pos := 0
	for i := 2; i < len(input)-1; i++ {
		if input[i] == byte('\r') && input[i+1] == byte('\n') {
			pos = i + 1
			break // Break if you only need the first occurrence
		} else {
			// check if valid int
			_, err := strconv.ParseInt(string(input[i]), 10, 64)
			if err != nil {
				return 0, 0, fmt.Errorf("not a valid integer input")
			}
		}
	}

	if pos == 0 {
		return 0, 0, fmt.Errorf("not a valid integer input")
	}

	end := pos - 1
	if start == end {
		return 0, start, nil
	}

	value, err := strconv.Atoi(string(input[start:end]))
	if err != nil {
		return 0, 0, err
	}
	return value, pos - 1, nil
}

// func decodeBoolean(input []byte) (bool, int, error) {
// 	return false, 0, nil
// }

// // func decodeArrays(input []byte) ([]interface{}, int, error) {

// // 	return []interface{}, 0, nil
// // }

func decodeSimpleString(input []byte) (string, int, error) {
	if len(input) < 3 {
		return "", 0, fmt.Errorf("not a valid string input. length < 3")
	}
	if input[0] != '+' {
		return "", 0, fmt.Errorf("not a valid string input. should start with +")
	}

	pos := 0
	for i := 1; i < len(input)-1; i++ {
		if input[i] == byte('\r') && input[i+1] == byte('\n') {
			pos = i + 1
			break // Break if you only need the first occurrence
		}
	}
	if pos == 0 {
		return "", 0, fmt.Errorf("not a valid string input")
	}
	if pos == 2 {
		return "", 1, nil
	}
	return string(input[1 : pos-1]), pos - 1, nil
}

func decodeBulkString(input []byte) (string, int, error) {
	if len(input) < 6 {
		return "", 0, fmt.Errorf("not a valid bulk string input. length < 6")
	}
	if input[0] != '$' {
		return "", 0, fmt.Errorf("not a valid bulk string input. should start with $")
	}
	pos := 0
	for i := 1; i < len(input)-1; i++ {
		if input[i] == byte('\r') && input[i+1] == byte('\n') {
			pos = i + 1
			break // Break if you only need the first occurrence
		} else {
			_, err := strconv.Atoi(string(input[i]))
			if err != nil {
				return "", 0, err
			}
		}
	}
	if pos == 0 {
		return "", 0, fmt.Errorf("not a valid bulk string input. should have length of string provided")
	}

	length, err := strconv.Atoi(string(input[1 : pos-1]))
	if err != nil {
		return "", 0, fmt.Errorf("not a valid bulk string input. should have length of string provided '%w'", err)
	}

	stringStart := pos + 1
	stringEnd := pos + length
	if stringEnd+2 != len(input)-1 {
		return "", 0, fmt.Errorf("not a valid bulk string input")
	}
	if input[stringEnd+1] != byte('\r') && input[stringEnd+2] != byte('\n') {
		return "", 0, fmt.Errorf("not a valid bulk string input. should end with CLRF: %w", err)
	}
	return string(input[stringStart : stringEnd+1]), stringEnd + 1, nil
}
