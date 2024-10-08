package core

import (
	"fmt"
	"net"
	"strings"
)

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

func ReadAndEval(data []byte) (string, error) {
	tokens, err := convertByteArrayToStringArray(data)
	if err != nil {
		return "", err
	}
	return ProcessCmd(&RedisCmd{Cmd: tokens[0], Args: tokens[1:]})
}

func ProcessCmd(cmd *RedisCmd) (string, error) {
	if cmd == nil {
		return "", fmt.Errorf("cmd cannot be nil")
	}
	switch strings.ToLower(cmd.Cmd) {
	case "ping":
		{
			if len(cmd.Args) == 0 {
				return "PONG", nil
			} else if len(cmd.Args) == 1 {
				return cmd.Args[0], nil
			} else {
				return "", fmt.Errorf("wrong number of arguments to PING cmd")
			}
		}
	default:
		return "hi client", nil
	}
}

// convertByteArrayToStringArray is helper function
// the input from the redis cli is always sent as array of string
func convertByteArrayToStringArray(data []byte) ([]string, error) {
	tokens, _, err := Decode(data)
	if err != nil {
		return []string{}, err
	}

	values := tokens.([]any)
	output := make([]string, len(values))
	for i := 0; i < len(values); i++ {
		output[i] = values[i].(string)
	}
	return output, nil
}
