package store

import "errors"

type Store struct {
	dataMap map[string]any
}

func NewStore() *Store {
	dataMap := make(map[string]any, 0)
	store := new(Store)
	store.dataMap = dataMap
	return store
}

func (s *Store) Set(key string, value any) {
	s.dataMap[key] = value
}

func (s *Store) Get(key string) (any, error) {
	if val, ok := s.dataMap[key]; !ok {
		return nil, errors.New("not found")
	} else {
		return val, nil
	}
}
