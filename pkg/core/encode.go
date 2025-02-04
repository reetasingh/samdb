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
	value := fmt.Sprintf(":%d\r\n", input)
	return []byte(value)
}

func EncodeError(err error) []byte {
	value := fmt.Sprintf("-%s\r\n", err.Error())
	return []byte(value)
}

func EncodeOne(input any) ([]byte, error) {
	switch input.(type) {
	case int64, int, int32:
		return EncodeInt(input.(int64)), nil
	case string:
		return EncodeString(input.(string), false), nil
	case error:
		return EncodeError(input.(error)), nil
	case []string:
		return EncodeArray(input.([]string))
	default:
		return []byte{}, fmt.Errorf("wrong type")
	}
}

func EncodeArray(input []string) ([]byte, error) {
	n := len(input)
	value := fmt.Sprintf("*%d\r\n", n)
	for _, item := range input {
		v, err := EncodeOne(item)
		if err != nil {
			return v, err
		}
		value = fmt.Sprintf("%s%s", value, string(v))
	}
	return []byte(value), nil
}
