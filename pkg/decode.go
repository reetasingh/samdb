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

// func decodeSimpleString(input []byte) (string, int, error) {

// 	return "", 0, nil
// }

// func decodeBulkString(input []byte) (string, int, error) {
// 	return "", 0, nil
// }
