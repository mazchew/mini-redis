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
	fmt.Println("=== before ===")
	fmt.Println(kv.store)
	kv.store[key] = value
	fmt.Println("=== after ===")
	fmt.Println(kv.store)
}

func (kv *KVStore) Get(key string) (string, error) {
	fmt.Println("=== getting ===")
	fmt.Println("=== key: %s ===", key)
	fmt.Println(kv.store)
	val, ok := kv.store[key]
	fmt.Println("==== got ====")
	fmt.Println(val)
	if !ok {
		return "", fmt.Errorf("key %s does not exist", key)
	}
	return val, nil
}
