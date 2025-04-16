package cmd

import (
	"fmt"
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

func ReadAndEvalSingleCmd(input []byte, dbStore store.DBStore) []byte {
	tokens, err := ReadStringTokens(input)
	if err != nil {
		return RespondForSingleCmd([]byte{}, err)
	}
	result, err := ProcessCmd(&RedisCmd{Cmd: tokens[0], Args: tokens[1:]}, dbStore)
	return RespondForSingleCmd(result, err)
}

func RespondForSingleCmd(response []byte, err error) []byte {
	if err != nil {
		return core.EncodeError(err)
	}
	return response
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
	case "bgrewriteaof":
		{
			return evalbgREWRITEAOF(dbStore)
		}
	default:
		return core.EncodeString("hi client", false), nil
	}
}

// ReadTokens is helper function
// the input from the redis cli is always an array of string
func ReadStringTokens(data []byte) ([]string, error) {
	// everytime we get byte array from REDIS
	tokens, err := core.Decode(data)
	if err != nil {
		return []string{}, err
	}
	return toArrayString(tokens)
}

func toArrayString(values []any) ([]string, error) {
	output := make([]string, 0)
	for i := 0; i < len(values); i++ {
		switch values[i].(type) {
		case []any:
			multipleString, err := toArrayString(values[i].([]any))
			if err != nil {
				return nil, err
			}
			output = append(output, multipleString...)
		case string:
			output = append(output, values[i].(string))
		default:
			return nil, fmt.Errorf("invalid type")
		}
	}
	return output, nil
}
