package cmd

import (
	"bytes"

	"github.com/reetasingh/samdb/pkg/core"
	"github.com/reetasingh/samdb/pkg/store"
)

func ProcessCmds(cmds []RedisCmd, dbStore store.DBStore) []byte {
	output := make([]byte, 0)
	outputBuffer := bytes.NewBuffer(output)
	for _, c := range cmds {
		result, err := ProcessCmd(&c, dbStore)
		if err != nil {
			outputBuffer.Write(core.EncodeError(err))
		}
		outputBuffer.Write(result)
	}
	return outputBuffer.Bytes()
}

func ReadMultipleCmdsTokens(data []byte) ([]RedisCmd, error) {
	// everytime we get byte array from REDIS
	values, err := core.Decode(data)
	if err != nil {
		return nil, err
	}
	cmds := make([]RedisCmd, 0)
	for _, v := range values {
		tokens, err := toArrayString(v.([]any))
		if err != nil {
			return nil, err
		}
		cmd := RedisCmd{Cmd: tokens[0], Args: tokens[1:]}
		cmds = append(cmds, cmd)
	}
	return cmds, nil
}

func ReadAndEvalMultipleCmds(data []byte, dbStore store.DBStore) []byte {
	cmds, err := ReadMultipleCmdsTokens(data)
	if err != nil {
		return Respond([]byte{}, err)
	}
	return ProcessCmds(cmds, dbStore)
}

func Respond(response []byte, err error) []byte {
	if err != nil {
		return core.EncodeError(err)
	}
	return response
}
