package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/reetasingh/samdb/pkg/core"
	"github.com/reetasingh/samdb/pkg/store"
)

// TODO: add unit tests
type RedisCmd struct {
	Cmd  string
	Args []string
}

func NILValue() []byte {
	return []byte("$-1\r\n")
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

func ReadAndEval(data []byte, dbStore store.DBStore) ([]byte, error) {
	tokens, err := convertByteArrayToStringArray(data)
	if err != nil {
		return []byte{}, err
	}
	return ProcessCmd(&RedisCmd{Cmd: tokens[0], Args: tokens[1:]}, dbStore)
}

func ProcessCmd(cmd *RedisCmd, dbStore store.DBStore) ([]byte, error) {
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
			return evalGet(cmd, dbStore)
		}
	case "set":
		{
			return evalSet(cmd, dbStore)
		}
	case "ttl":
		{
			return evalTTL(cmd, dbStore)
		}
	case "del":
		{
			return evalDelete(cmd, dbStore)
		}
	case "expire":
		{
			return evalExpire(cmd, dbStore)
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

func evalGet(cmd *RedisCmd, dbStore store.DBStore) ([]byte, error) {
	if len(cmd.Args) > 1 || len(cmd.Args) < 1 {
		return []byte{}, fmt.Errorf("wrong number of arguments to GET command")
	}
	if val, err := dbStore.Get(cmd.Args[0]); err != nil {
		// nil
		if errors.Is(err, store.KeyNotFound{}) {
			return []byte("$-1\r\n"), nil
		}
		return []byte{}, err

	} else {
		return core.EncodeString(val.(string), true), nil
	}
}

func evalSet(cmd *RedisCmd, dbStore store.DBStore) ([]byte, error) {
	n := len(cmd.Args)
	if n < 2 || n > 4 || n == 3 {
		return []byte{}, fmt.Errorf("wrong number of arguments to SET command")
	}
	key := cmd.Args[0]
	value := cmd.Args[1]
	ttlSeconds := int64(-1)
	var err error
	for i := 2; i < n; i = i + 1 {
		if strings.ToLower(cmd.Args[i]) == "ex" {
			if i == n-1 {
				return []byte{}, fmt.Errorf("wrong number of arguments to SET command")
			}
			ttlSeconds, err = strconv.ParseInt(cmd.Args[i+1], 10, 64)
			if err != nil {
				return []byte{}, fmt.Errorf("wrong value for %s", cmd.Args[i+1])
			}
		}
	}
	dbStore.Set(key, value, ttlSeconds)
	return core.EncodeString("OK", false), nil
}

func evalTTL(cmd *RedisCmd, dbStore store.DBStore) ([]byte, error) {
	if len(cmd.Args) > 1 || len(cmd.Args) == 0 {
		return []byte{}, fmt.Errorf("wrong number of arguments to TTL command")
	}
	key := cmd.Args[0]
	if val, err := dbStore.GetTTL(key); err != nil {
		// nil
		return []byte("$-1\r\n"), nil
	} else {
		return core.EncodeInt(val), nil
	}
}

func evalDelete(cmd *RedisCmd, store store.DBStore) ([]byte, error) {
	if len(cmd.Args) == 0 {
		return core.EncodeInt(0), fmt.Errorf("wrong number of arguments to Delete command")
	}
	count := int64(0)
	for _, key := range cmd.Args {
		if ok := store.Delete(key); ok {
			count = count + 1
		}
	}

	return core.EncodeInt(count), nil
}

func evalExpire(cmd *RedisCmd, store store.DBStore) ([]byte, error) {
	if len(cmd.Args) < 2 {
		return []byte{}, fmt.Errorf("wrong number of arguments to Expire command")
	}
	count := int64(0)
	key := cmd.Args[0]
	ttlSeconds, err := strconv.ParseInt(cmd.Args[1], 10, 64)
	if err != nil {
		return core.EncodeInt(count), fmt.Errorf("wrong value for %s", cmd.Args[0])
	}
	if ok := store.SetTTL(key, ttlSeconds); ok {
		count = 1
	}

	return core.EncodeInt(count), nil
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
