package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"

	"github.com/reetasingh/samdb/pkg/core"
	"github.com/reetasingh/samdb/pkg/store"
)

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
	if keyValue, err := dbStore.Get(cmd.Args[0]); err != nil {
		// nil
		if errors.Is(err, store.KeyNotFound{}) {
			return []byte("$-1\r\n"), nil
		}
		return []byte{}, err
	} else {
		returnValue := keyValue.Value
		// TODO check keyValue.typeEncoding for return and call encode accordingly
		return core.EncodeString(returnValue.(string), true), nil
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
	err = dbStore.Set(key, value, ttlSeconds)
	if err != nil {
		return []byte{}, err
	}
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

func evalbgREWRITEAOF(dbStore store.DBStore) ([]byte, error) {
	file, err := os.OpenFile("samdb.aof", os.O_CREATE|os.O_APPEND, fs.ModeAppend)
	defer file.Close()
	if err != nil {
		return []byte{}, err
	}
	all := dbStore.GetAll()
	for k, v := range all {
		err := dumpKey(file, k, v)
		if err != nil {
			return []byte{}, err
		}
	}
	return core.EncodeString("OK", false), nil
}

func evalINCR(cmd *RedisCmd, dbStore store.DBStore) ([]byte, error) {
	if len(cmd.Args) > 1 || len(cmd.Args) < 1 {
		return []byte{}, fmt.Errorf("wrong number of arguments to INCR command")
	}
	if keyValue, err := dbStore.Get(cmd.Args[0]); err != nil {
		// nil
		if errors.Is(err, store.KeyNotFound{}) {
			// TODO store in DB for next increment on same key
			return core.EncodeInt(0), nil
		}
		return []byte{}, err
	} else {
		returnValue := keyValue.Value
		encodingValue := keyValue.TypeEncoding & 15
		if encodingValue == core.OBJ_INTEGER_ENCODING {
			intValue, err := strconv.Atoi(returnValue.(string))
			if err != nil {
				return []byte{}, err
			}
			return core.EncodeInt(int64(intValue)), nil
		}
		// TODO return error since incr only supports integer
		return core.EncodeString(returnValue.(string), true), nil
	}
}

func dumpKey(fp *os.File, key string, value any) error {
	cmd := fmt.Sprintf("SET %s %s", key, value)
	tokens := strings.Split(cmd, " ")
	content, err := core.EncodeArray(tokens)
	if err != nil {
		return err
	}
	fp.Write(content)
	return nil
}
