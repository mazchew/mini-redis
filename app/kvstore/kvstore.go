package kvstore

import (
	"fmt"
	"time"
)

type KVStore struct {
	store map[string]*Value
}

func NewKVStore() *KVStore {
	return &KVStore{
		store: make(map[string]*Value),
	}
}

func (kv *KVStore) Set(key string, value interface{}, ttl int64) {
	kv.store[key] = NewValue(value, ttl)
}

func (kv *KVStore) Get(key string) (*Value, error) {
	val, exists := kv.store[key]
	if !exists {
		return nil, fmt.Errorf("key %s does not exist", key)
	}
	currentTime := time.Now().UnixNano() / 1e6
	if val.Ttl > 0 && currentTime > val.SetAt + val.Ttl {
		delete(kv.store, key)
		return nil, fmt.Errorf("key %s has expired", key)
	}
	return val, nil
}
