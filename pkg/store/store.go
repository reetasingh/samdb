package store

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/reetasingh/samdb/pkg/core"
)

type TTLExpiredErr struct{}
type KeyNotFound struct{}

func (t TTLExpiredErr) Error() string {
	return "TTL expired"
}

func (t KeyNotFound) Error() string {
	return "key not found"
}

type DBStore interface {
	Set(key string, value any, ttlSeconds int64) error
	Get(key string) (*Data, error)
	Delete(key string) bool
	GetTTL(key string) (int64, error)
	SetTTL(key string, ttlSeconds int64) bool
	CleanupExpiredKeys()
	GetAll() map[string]any
}

type DBStoreImpl struct {
	dataMap  map[string]Data
	keyLimit int
}

type Data struct {
	Value any
	// most significant 4 bits for type
	// least significant 4 bits for encoding
	TypeEncoding uint8
	TTLSeconds   int64
}

func NewDBStore(keyLimit int) *DBStoreImpl {
	dataMap := make(map[string]Data, 0)
	store := new(DBStoreImpl)
	store.dataMap = dataMap
	store.keyLimit = keyLimit
	return store
}

// randomly delete 1 element
func (s *DBStoreImpl) evict() {
	for k := range s.dataMap {
		// just delete 1 element
		s.Delete(k)
		return
	}
}

func (s *DBStoreImpl) Set(key string, value any, ttlSeconds int64) error {
	if len(s.dataMap) >= s.keyLimit {
		s.evict()
	}
	typeEncoding := uint8(core.OBJ_STRING_TYPE)
	switch value.(type) {
	case string:
		typeEncoding = core.OBJ_STRING_TYPE
		typeEncoding = typeEncoding << 4
		_, err := strconv.Atoi(value.(string))
		if err == nil {
			typeEncoding = typeEncoding | core.OBJ_INTEGER_ENCODING
		} else {
			typeEncoding = typeEncoding | core.OBJ_STRING_ENCODING
		}
	default:
		return fmt.Errorf("type not supported")
	}
	data := Data{Value: value, TypeEncoding: typeEncoding}
	data.TTLSeconds = ttlSeconds
	if data.TTLSeconds != -1 {
		data.TTLSeconds = time.Now().Unix() + data.TTLSeconds
	} else {
		data.TTLSeconds = -1
	}
	s.dataMap[key] = data
	return nil
}

func (s *DBStoreImpl) Get(key string) (*Data, error) {
	if val, ok := s.dataMap[key]; !ok {
		return nil, KeyNotFound{}
	} else {
		_, err := s.GetTTL(key)
		if errors.Is(err, TTLExpiredErr{}) {
			return nil, KeyNotFound{}
		}
		return &val, nil
	}
}

func (s *DBStoreImpl) Delete(key string) bool {
	if _, ok := s.dataMap[key]; ok {
		delete(s.dataMap, key)
		return true
	}
	return false
}

func (s *DBStoreImpl) GetTTL(key string) (int64, error) {
	if val, ok := s.dataMap[key]; !ok {
		return -1, errors.New("not found")
	} else {
		if val.TTLSeconds == -1 {
			// no TTL set
			return -1, nil
		}
		timeDiffSeconds := val.TTLSeconds - time.Now().Unix()
		if timeDiffSeconds > 0 {
			return timeDiffSeconds, nil
		} else {
			return -1, TTLExpiredErr{}
		}
	}
}

func (s *DBStoreImpl) SetTTL(key string, ttlSeconds int64) bool {
	if data, ok := s.dataMap[key]; ok {
		newttlSeconds := time.Now().Unix() + ttlSeconds
		newData := Data{data.Value, data.TypeEncoding, newttlSeconds}
		s.dataMap[key] = newData
		return true
	}
	return false
}

/* Keys are never deleted automatically when their TTL expires;
they remain in the system unless explicitly removed by the user.
Periodically, we randomly select 20% of the keys and delete those that are expired.
This process continues until fewer than 80% of the randomly selected 20% keys are expired.
This is a best effort cleanup process*/
func (s *DBStoreImpl) CleanupExpiredKeys() {
	fmt.Println("cleaning up expired keys")
	totalCount := len(s.dataMap)
	// 20 % totalCount
	randomKeysCount := int(float32(totalCount) * 0.2)
	fmt.Println(randomKeysCount)
	threshholdCount := int(float32(randomKeysCount) * 0.8)
	c := 0
	allKeys := []string{}
	// slice of all keys
	for k := range s.dataMap {
		allKeys = append(allKeys, k)
	}

	src := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(src)
	// randomly select 20% of the keys
	for i := 0; i < randomKeysCount; i++ {
		randomPos := randomGenerator.Intn(randomKeysCount)
		k := allKeys[randomPos]
		fmt.Printf("\nrandomly selected %s", k)
		_, err := s.GetTTL(k)
		if errors.Is(err, TTLExpiredErr{}) {
			s.Delete(k)
			c = c + 1
		}
	}
	fmt.Printf("\ncleaned up %d keys:", c)

	if c > 0 && c == threshholdCount {
		// 80% of expired keys from the batch of 20% means there could be more keys which are expired
		// so we repeat again
		s.CleanupExpiredKeys()
	}
}

func (s *DBStoreImpl) GetAll() map[string]any {
	result := make(map[string]any, len(s.dataMap))
	for k, v := range s.dataMap {
		result[k] = v.Value
	}

	return result
}
