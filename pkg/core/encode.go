package core

import "fmt"

func EncodeString(input string) []byte {
	value := fmt.Sprintf("$%d\r\n%s\r\n", len(input), input)
	return []byte(value)
}

func EncodeError(err error) []byte {
	value := fmt.Sprintf("-%s\r\n", err.Error())
	return []byte(value)
}
