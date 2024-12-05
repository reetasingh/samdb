package store

import (
	"errors"
	"fmt"
	"time"
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
	Set(key string, value any, ttlSeconds int64)
	Get(key string) (any, error)
	Delete(key string) bool
	GetTTL(key string) (int64, error)
	SetTTL(key string, ttlSeconds int64) bool
}

type DBStoreImpl struct {
	dataMap map[string]Data
}

type Data struct {
	value      any
	ttlSeconds int64
}

func NewDBStore() *DBStoreImpl {
	dataMap := make(map[string]Data, 0)
	store := new(DBStoreImpl)
	store.dataMap = dataMap
	return store
}

func (s *DBStoreImpl) Set(key string, value any, ttlSeconds int64) {
	data := Data{value: value}
	data.ttlSeconds = ttlSeconds
	if data.ttlSeconds != -1 {
		data.ttlSeconds = time.Now().Unix() + data.ttlSeconds
	} else {
		data.ttlSeconds = -1
	}
	s.dataMap[key] = data
}

func (s *DBStoreImpl) Get(key string) (any, error) {
	if val, ok := s.dataMap[key]; !ok {
		return nil, KeyNotFound{}
	} else {
		_, err := s.GetTTL(key)
		if errors.Is(err, TTLExpiredErr{}) {
			return nil, KeyNotFound{}
		}
		return val.value, nil
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
		if val.ttlSeconds == -1 {
			return -1, nil
		}
		timeDiffSeconds := val.ttlSeconds - time.Now().Unix()
		fmt.Println(timeDiffSeconds)
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
		newData := Data{data.value, newttlSeconds}
		s.dataMap[key] = newData
		return true
	}
	return false
}
