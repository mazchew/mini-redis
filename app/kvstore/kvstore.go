package kvstore

import (
	"fmt"
	"time"

	"github.com/codecrafters-io/redis-starter-go/app/utils"
)

type KVStore struct {
	Store map[string]*Value
}

func NewKVStore() *KVStore {
	return &KVStore{
		Store: make(map[string]*Value),
	}
}

func (kv *KVStore) Set(key string, value interface{}, ttl ...int64) {
	var ttlVal int64
	if len(ttl) > 0 {
		ttlVal = ttl[0]
	} else {
		ttlVal = -1
	}
	kv.Store[key] = NewValueWithTtl(value, ttlVal)
}

func (kv *KVStore) SetWithExpiry(key string, value interface{}, expiryTime int64) {
	if time.Now().UnixMilli() < expiryTime {
		kv.Store[key] = NewValueWithExpiry(value, expiryTime)
	} else if expiryTime == -1 {
		kv.Store[key] = NewValueWithExpiry(value, expiryTime)
	}
}

func (kv *KVStore) Get(key string) (*Value, error) {
	val, exists := kv.Store[key]
	if !exists {
		return nil, fmt.Errorf("key %s does not exist", key)
	}

	if val.ExpiryTime > 0 && time.Now().UnixMilli() > val.ExpiryTime {
		delete(kv.Store, key)
		return nil, fmt.Errorf("key %s has expired", key)
	}
	return val, nil
}

func (kv *KVStore) WriteToCacheFromRDBFile(filepath string) {
	data := utils.ParseFile(filepath)

	for _, cacheRecord := range data {
		key := cacheRecord.Key
		value := cacheRecord.Value
		expiryTime := cacheRecord.ExpiryTime
		kv.SetWithExpiry(key, value, expiryTime)
	}
}
