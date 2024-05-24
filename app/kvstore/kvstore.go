package kvstore

import (
	"fmt"
	"time"

	"github.com/codecrafters-io/redis-starter-go/app/utils"
)

type KVStore struct {
	store map[string]*Value
}

func NewKVStore() *KVStore {
	return &KVStore{
		store: make(map[string]*Value),
	}
}

func (kv *KVStore) Set(key string, value interface{}, ttl ...int64) {
	var ttlVal int64
	if len(ttl) > 0 {
		ttlVal = ttl[0]
	} else {
		ttlVal = -1
	}
	kv.store[key] = NewValue(value, ttlVal)
}

func (kv *KVStore) Get(key string) (*Value, error) {
	val, exists := kv.store[key]
	if !exists {
		return nil, fmt.Errorf("key %s does not exist", key)
	}
	currentTime := time.Now().UnixNano() / 1e6
	if val.Ttl > 0 && currentTime > val.SetAt+val.Ttl {
		delete(kv.store, key)
		return nil, fmt.Errorf("key %s has expired", key)
	}
	return val, nil
}

func (kv *KVStore) WriteToCacheFromRDBFile(filepath string) {
	data := utils.ParseFile(filepath)

	for _, cacheRecord := range data {
		key := cacheRecord.Key
		value := cacheRecord.Value
		ttl := cacheRecord.Ttl
		kv.Set(key, value, ttl)
	}
}
