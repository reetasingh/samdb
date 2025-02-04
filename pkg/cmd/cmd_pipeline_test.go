package cmd

// import (
// 	"testing"

// 	"github.com/reetasingh/samdb/mocks"
// 	"github.com/stretchr/testify/assert"
// )

// func TestReadMultipleCmdsTokens_PingCommand(t *testing.T) {
// 	args := struct {
// 		data []byte
// 	}{
// 		data: []byte("*1\r\n$4\r\nPING\r\n"),
// 	}

// 	got, err := ReadMultipleCmdsTokens(args.data)
// 	if err != nil {
// 		t.Errorf("ReadMultipleCmdsTokens() error = %v, wantErr %v", err, false)
// 	}
// 	assert.Equal(t, got[0].Cmd, "PING")
// 	assert.Equal(t, len(got[0].Args), 0)
// }

// func TestReadMultipleCmdsTokens_PingCommandWithSet(t *testing.T) {
// 	args := struct {
// 		data []byte
// 	}{
// 		data: []byte("*1\r\n$4\r\nPING\r\n*3\r\n$3\r\nSET\r\n$4\r\nkey1\r\n$6\r\nvalue1\r\n"),
// 	}

// 	got, err := ReadMultipleCmdsTokens(args.data)
// 	if err != nil {
// 		t.Errorf("ReadMultipleCmdsTokens() error = %v, wantErr %v", err, false)
// 	}
// 	assert.Equal(t, got[0].Cmd, "PING")
// 	assert.Equal(t, len(got[0].Args), 0)

// 	assert.Equal(t, got[1].Cmd, "SET")
// 	assert.Equal(t, got[1].Args[0], "key1")
// 	assert.Equal(t, got[1].Args[1], "value1")
// }

// func TestReadMultipleCmdsTokens_SetGetTtlCommands(t *testing.T) {
// 	args := struct {
// 		data []byte
// 	}{
// 		data: []byte("*3\r\n$3\r\nSET\r\n$4\r\nkey1\r\n$6\r\nvalue1\r\n*2\r\n$3\r\nGET\r\n$4\r\nkey1\r\n*2\r\n$3\r\nTTL\r\n$4\r\nkey1\r\n"),
// 	}

// 	got, err := ReadMultipleCmdsTokens(args.data)
// 	if err != nil {
// 		t.Errorf("ReadMultipleCmdsTokens() error = %v, wantErr %v", err, false)
// 	}
// 	assert.Equal(t, got[0].Cmd, "SET")
// 	assert.Equal(t, got[0].Args[0], "key1")
// 	assert.Equal(t, got[0].Args[1], "value1")

// 	assert.Equal(t, got[1].Cmd, "GET")
// 	assert.Equal(t, got[1].Args[0], "key1")

// 	assert.Equal(t, got[2].Cmd, "TTL")
// 	assert.Equal(t, got[2].Args[0], "key1")
// }

// func createMockDB(key1 string, value1 string) *mocks.DBStore {
// 	dbMock := &mocks.DBStore{}
// 	dbMock.On("Set", key1, value1, int64(-1)).Return(nil)
// 	dbMock.On("Get", key1).Return(value1, nil)
// 	dbMock.On("TTL", key1).Return(value1)
// 	return dbMock
// }

// func TestProcessCmds(t *testing.T) {
// 	args := struct {
// 		data []byte
// 	}{
// 		data: []byte("*1\r\n$4\r\nPING\r\n"),
// 	}

// 	got, err := ReadMultipleCmdsTokens(args.data)
// 	if err != nil {
// 		t.Errorf("ReadMultipleCmdsTokens() error = %v, wantErr %v", err, false)
// 	}
// 	// assert.Equal(t, got[0].Cmd, "SET")
// 	// assert.Equal(t, got[0].Args[0], "key1")
// 	// assert.Equal(t, got[0].Args[1], "value1")

// 	// assert.Equal(t, got[1].Cmd, "GET")
// 	// assert.Equal(t, got[1].Args[0], "key1")

// 	// assert.Equal(t, got[2].Cmd, "TTL")
// 	// assert.Equal(t, got[2].Args[0], "key1")

// 	result := ProcessCmds(got, createMockDB("key1", "value1"))
// 	assert.Equal(t, string(result), "$4\r\nPONG\r\n")

// }
