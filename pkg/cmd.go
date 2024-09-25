package core

import "net"

type RedisCmd struct {
	Cmd  string
	Args []string
}

func ReadCmd(conn net.Conn) (*RedisCmd, error) {
	data := make([]byte, 1024)
	_, err := conn.Read(data)
	if err != nil {
		return nil, err
	}
	tokens, err := convertByteArrayToStringArray(data)
	if err != nil {
		return nil, err
	}
	cmd := RedisCmd{Cmd: tokens[0], Args: tokens[1:]}
	return &cmd, nil
}

// convertByteArrayToStringArray is helper function
// the input from the redis cli is always sent as array of string
func convertByteArrayToStringArray(data []byte) ([]string, error) {
	tokens, _, err := decode(data)
	if err != nil {
		return []string{}, err
	}
	n := len(data)

	values := tokens.([]interface{})
	output := make([]string, n)
	for i := 0; i < n; i++ {
		output[i] = values[i].(string)
	}
	return output, nil
}
