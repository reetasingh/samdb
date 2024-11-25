package core

import "fmt"

func EncodeString(input string, bulk bool) []byte {
	var value string
	if bulk {
		value = fmt.Sprintf("$%d\r\n%s\r\n", len(input), input)
	} else {
		value = fmt.Sprintf("+%s\r\n", input)
	}

	return []byte(value)
}

func EncodeInt(input int64) []byte {
	var value string
	value = fmt.Sprintf(":%d\r\n", input)
	fmt.Println(value)
	return []byte(value)
}

func EncodeError(err error) []byte {
	value := fmt.Sprintf("-%s\r\n", err.Error())
	return []byte(value)
}
