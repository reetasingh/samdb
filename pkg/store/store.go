package store

import (
	"errors"
	"fmt"
	"time"
)

type TTLExpiredErr struct{}

func (t TTLExpiredErr) Error() string {
	return "TTL expired"
}

type Store struct {
	dataMap map[string]Data
}

type Data struct {
	value      any
	ttlSeconds int64
}

func NewStore() *Store {
	dataMap := make(map[string]Data, 0)
	store := new(Store)
	store.dataMap = dataMap
	return store
}

func (s *Store) Set(key string, value any, ttlSeconds int64) {
	data := Data{value: value}
	data.ttlSeconds = ttlSeconds
	if data.ttlSeconds != -1 {
		data.ttlSeconds = time.Now().Unix() + data.ttlSeconds
	} else {
		data.ttlSeconds = -1
	}
	s.dataMap[key] = data
}

func (s *Store) Get(key string) (any, error) {
	if val, ok := s.dataMap[key]; !ok {
		return nil, errors.New("not found")
	} else {
		_, err := s.GetTTL(key)
		if errors.Is(err, TTLExpiredErr{}) {
			return nil, errors.New("not found")
		}
		return val.value, nil
	}
}

func (s *Store) GetTTL(key string) (int64, error) {
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
