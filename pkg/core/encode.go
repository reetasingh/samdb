package core

import "fmt"

func EncodeString(input string) []byte {
	value := fmt.Sprintf("$%d\r\n%s\r\n", len(input), input)
	return []byte(value)
}
