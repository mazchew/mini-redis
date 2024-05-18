package kvstore

import "fmt"

type KVStore struct {
	store map[string]string
}

func NewKVStore() *KVStore {
	return &KVStore{
		store: make(map[string]string),
	}
}

func (kv *KVStore) Set(key string, value string) {
	kv.store[key] = value
}

func (kv *KVStore) Get(key string) (string, error) {
	val, ok := kv.store[key]
	if !ok {
		return "", fmt.Errorf("key %s does not exist", key)
	}
	return val, nil
}
