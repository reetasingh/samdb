package cmd

// import (
// 	"errors"
// 	"reflect"
// 	"testing"

// 	"github.com/reetasingh/samdb/mocks"
// 	"github.com/reetasingh/samdb/pkg/core"
// 	"github.com/reetasingh/samdb/pkg/store"
// )

// func Test_evalPing(t *testing.T) {
// 	type args struct {
// 		cmd *RedisCmd
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    []byte
// 		wantErr bool
// 	}{
// 		{
// 			name: "one argument",
// 			args: args{
// 				cmd: &RedisCmd{
// 					Args: []string{"hello"},
// 				},
// 			},
// 			want:    []byte(core.EncodeString("hello", true)),
// 			wantErr: false,
// 		},
// 		{
// 			name: "two arguments",
// 			args: args{
// 				cmd: &RedisCmd{
// 					Args: []string{"hello", "hi"},
// 				},
// 			},
// 			want:    []byte{},
// 			wantErr: true,
// 		},
// 		{
// 			name: "no arguments",
// 			args: args{
// 				cmd: &RedisCmd{
// 					Args: []string{},
// 				},
// 			},
// 			want:    []byte(core.EncodeString("PONG", true)),
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := evalPing(tt.args.cmd)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("evalPing() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("evalPing() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_evalGet(t *testing.T) {
// 	type args struct {
// 		cmd     *RedisCmd
// 		dbStore store.DBStore
// 	}

// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    []byte
// 		wantErr bool
// 	}{
// 		{
// 			name: "existing key",
// 			args: args{
// 				cmd: &RedisCmd{
// 					Args: []string{"a"},
// 				},
// 				dbStore: createMockDBForGet("a", "b", nil),
// 			},
// 			want:    []byte(core.EncodeString("b", true)),
// 			wantErr: false,
// 		},
// 		{
// 			name: "not existing key",
// 			args: args{
// 				cmd: &RedisCmd{
// 					Args: []string{"a"},
// 				},
// 				dbStore: createMockDBForGet("a", "", store.KeyNotFound{}),
// 			},
// 			want:    []byte(NILValue()),
// 			wantErr: false,
// 		},
// 		{
// 			name: "error",
// 			args: args{
// 				cmd: &RedisCmd{
// 					Args: []string{"a"},
// 				},
// 				dbStore: createMockDBForGet("a", nil, errors.New("failure")),
// 			},
// 			want:    []byte{},
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := evalGet(tt.args.cmd, tt.args.dbStore)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("evalGet() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("evalGet() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_evalSet(t *testing.T) {
// 	type args struct {
// 		cmd     *RedisCmd
// 		dbStore store.DBStore
// 	}

// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    []byte
// 		wantErr bool
// 	}{
// 		{
// 			name: "< 2 arguments",
// 			args: args{
// 				cmd: &RedisCmd{
// 					Args: []string{"a"},
// 				},
// 				dbStore: createMockDBForGet("a", "b", nil),
// 			},
// 			want:    []byte{},
// 			wantErr: true,
// 		},
// 		{
// 			name: "happy path",
// 			args: args{
// 				cmd: &RedisCmd{
// 					Args: []string{"a", "b", "ex", "100"},
// 				},
// 				dbStore: createMockDBForSet("a", "b", 100),
// 			},
// 			want:    []byte(core.EncodeString("OK", false)),
// 			wantErr: false,
// 		},
// 		{
// 			name: "wrong value for expiry seconds",
// 			args: args{
// 				cmd: &RedisCmd{
// 					Args: []string{"a", "b", "ex", "1.2"},
// 				},
// 				dbStore: createMockDBForSet("a", "b", 100),
// 			},
// 			want:    []byte{},
// 			wantErr: true,
// 		},
// 		{
// 			name: "wrong numbers of args",
// 			args: args{
// 				cmd: &RedisCmd{
// 					Args: []string{"a", "b", "ex"},
// 				},
// 				dbStore: createMockDBForSet("a", "b", -1),
// 			},
// 			want:    []byte{},
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := evalSet(tt.args.cmd, tt.args.dbStore)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("evalSet() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("evalSet() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_evalDelete(t *testing.T) {
// 	type args struct {
// 		cmd   *RedisCmd
// 		store store.DBStore
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    []byte
// 		wantErr bool
// 	}{
// 		{
// 			name: "happy path",
// 			args: args{
// 				cmd: &RedisCmd{
// 					Args: []string{"key1", "key2", "key3"},
// 				},
// 				store: createMockDBForDelete([]string{"key1", "key2", "key3"}, []string{}),
// 			},
// 			want: core.EncodeInt(3),
// 		},
// 		{
// 			name: "one key missing",
// 			args: args{
// 				cmd: &RedisCmd{
// 					Args: []string{"key1", "key2", "key3"},
// 				},
// 				store: createMockDBForDelete([]string{"key1", "key2"}, []string{"key3"}),
// 			},
// 			want: core.EncodeInt(2),
// 		},
// 		{
// 			name: "wrong number of args",
// 			args: args{
// 				cmd: &RedisCmd{
// 					Args: []string{},
// 				},
// 			},
// 			want:    core.EncodeInt(0),
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := evalDelete(tt.args.cmd, tt.args.store)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("evalDelete() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("evalDelete() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func createMockDBForGet(key string, b any, err error) *mocks.DBStore {
// 	dbMock := &mocks.DBStore{}
// 	dbMock.On("Get", key).Return(b, err)
// 	return dbMock
// }

// func createMockDBForSet(key string, value any, ttlSeconds int64) *mocks.DBStore {
// 	dbMock := &mocks.DBStore{}
// 	dbMock.On("Set", key, value, ttlSeconds).Return(nil)
// 	return dbMock
// }

// func createMockDBForDelete(keysPresent []string, keysAbsent []string) *mocks.DBStore {
// 	dbMock := &mocks.DBStore{}
// 	for _, key := range keysPresent {
// 		dbMock.On("Delete", key).Return(true)
// 	}
// 	for _, key := range keysAbsent {
// 		dbMock.On("Delete", key).Return(false)
// 	}
// 	return dbMock
// }

// func createMockDBForGetTTL(key string, b any, err error) *mocks.DBStore {
// 	dbMock := &mocks.DBStore{}
// 	dbMock.On("GetTTL", key).Return(b, err)
// 	return dbMock
// }

// func createMockDBForSetTTL(key string, b int64, returnValue bool) *mocks.DBStore {
// 	dbMock := &mocks.DBStore{}
// 	dbMock.On("SetTTL", key, b).Return(returnValue)
// 	return dbMock
// }

// func Test_evalTTL(t *testing.T) {
// 	type args struct {
// 		cmd     *RedisCmd
// 		dbStore store.DBStore
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    []byte
// 		wantErr bool
// 	}{
// 		{
// 			name: "TTL expired",
// 			args: args{
// 				cmd: &RedisCmd{
// 					Args: []string{"a"},
// 				},
// 				dbStore: createMockDBForGetTTL("a", int64(-1), store.TTLExpiredErr{}),
// 			},
// 			want:    []byte("$-1\r\n"),
// 			wantErr: false,
// 		},
// 		{
// 			name: "valid",
// 			args: args{
// 				cmd: &RedisCmd{
// 					Args: []string{"a"},
// 				},
// 				dbStore: createMockDBForGetTTL("a", int64(10), nil),
// 			},

// 			want:    []byte(core.EncodeInt(int64(10))),
// 			wantErr: false,
// 		},
// 		{
// 			name: "key not found",
// 			args: args{
// 				cmd: &RedisCmd{
// 					Args: []string{"a"},
// 				},
// 				dbStore: createMockDBForGetTTL("a", int64(0), store.KeyNotFound{}),
// 			},
// 			want: []byte("$-1\r\n"),

// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := evalTTL(tt.args.cmd, tt.args.dbStore)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("evalTTL() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("evalTTL() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_evalExpire(t *testing.T) {
// 	type args struct {
// 		cmd   *RedisCmd
// 		store store.DBStore
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    []byte
// 		wantErr bool
// 	}{
// 		{
// 			name: "set invalid format of TTL",
// 			args: args{
// 				cmd: &RedisCmd{
// 					Args: []string{"a", "invalid ttl format"},
// 				},
// 			},
// 			want:    []byte(core.EncodeInt(int64(0))),
// 			wantErr: true,
// 		},
// 		{
// 			name: "happy path",
// 			args: args{
// 				cmd: &RedisCmd{
// 					Args: []string{"a", "10"},
// 				},
// 				store: createMockDBForSetTTL("a", 10, true),
// 			},
// 			want:    []byte(core.EncodeInt(int64(1))),
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := evalExpire(tt.args.cmd, tt.args.store)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("evalExpire() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("evalExpire() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
