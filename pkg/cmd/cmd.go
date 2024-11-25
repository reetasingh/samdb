package cmd

import (
	"fmt"
	"samdb/pkg/core"
	"samdb/pkg/store"
	"strconv"
	"strings"
)

type RedisCmd struct {
	Cmd  string
	Args []string
}

// func ReadCmd(conn net.Conn) (*RedisCmd, error) {
// 	data := make([]byte, 1024)
// 	_, err := conn.Read(data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	tokens, err := convertByteArrayToStringArray(data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	cmd := RedisCmd{Cmd: tokens[0], Args: tokens[1:]}
// 	return &cmd, nil
// }

func ReadAndEval(data []byte, store *store.Store) ([]byte, error) {
	tokens, err := convertByteArrayToStringArray(data)
	if err != nil {
		return []byte{}, err
	}
	return ProcessCmd(&RedisCmd{Cmd: tokens[0], Args: tokens[1:]}, store)
}

func ProcessCmd(cmd *RedisCmd, store *store.Store) ([]byte, error) {
	if cmd == nil {
		return []byte{}, fmt.Errorf("cmd cannot be nil")
	}
	switch strings.ToLower(cmd.Cmd) {
	case "ping":
		{
			return evalPing(cmd)
		}
	case "get":
		{
			return evalGet(cmd, store)
		}
	case "set":
		{
			return evalSet(cmd, store)
		}
	case "ttl":
		{
			return evalTTL(cmd, store)
		}
	default:
		return core.EncodeString("hi client", false), nil
	}
}

func evalPing(cmd *RedisCmd) ([]byte, error) {
	if len(cmd.Args) == 0 {
		return core.EncodeString("PONG", true), nil
	} else if len(cmd.Args) == 1 {
		return core.EncodeString(cmd.Args[0], true), nil
	} else {
		return []byte{}, fmt.Errorf("wrong number of arguments to PING cmd")
	}
}

func evalGet(cmd *RedisCmd, store *store.Store) ([]byte, error) {
	if len(cmd.Args) > 1 || len(cmd.Args) < 1 {
		return []byte{}, fmt.Errorf("wrong number of arguments to GET command")
	}
	if val, err := store.Get(cmd.Args[0]); err != nil {
		return []byte("$-1\r\n"), nil
	} else {
		return core.EncodeString(val.(string), true), nil
	}
}

func evalSet(cmd *RedisCmd, store *store.Store) ([]byte, error) {
	n := len(cmd.Args)
	if n < 2 {
		return []byte{}, fmt.Errorf("wrong number of arguments to SET command")
	}
	key := cmd.Args[0]
	value := cmd.Args[1]
	ttlSeconds := int64(-1)
	var err error
	for i := 2; i < n; i = i + 1 {
		if strings.ToLower(cmd.Args[i]) == "ex" {
			if i+1 == n {
				return []byte{}, fmt.Errorf("wrong number of arguments to SET command")
			}
			ttlSeconds, err = strconv.ParseInt(cmd.Args[i+1], 10, 64)
			if err != nil {
				return []byte{}, fmt.Errorf("wrong value for %s", cmd.Args[i+1])
			}
		}
	}
	store.Set(key, value, ttlSeconds)
	return core.EncodeString("OK", false), nil
}

func evalTTL(cmd *RedisCmd, store *store.Store) ([]byte, error) {
	if len(cmd.Args) > 1 || len(cmd.Args) < 1 {
		return []byte{}, fmt.Errorf("wrong number of arguments to TTL command")
	}
	key := cmd.Args[0]
	if val, err := store.GetTTL(key); err != nil {
		return []byte("$-1\r\n"), nil
	} else {
		return core.EncodeInt(val), nil
	}
}

// convertByteArrayToStringArray is helper function
// the input from the redis cli is always sent as array of string
func convertByteArrayToStringArray(data []byte) ([]string, error) {
	tokens, _, err := core.Decode(data)
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
