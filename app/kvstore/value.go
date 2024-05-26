package kvstore

import "time"

type Value struct {
	Val        interface{}
	ExpiryTime int64
}

func NewValueWithTtl(val interface{}, ttl int64) *Value {
	if ttl == -1 {
		return &Value{
			Val:        val,
			ExpiryTime: (time.Now().UnixNano() / 1e6) + ttl,
		}
	} else {
		return &Value{
			Val:        val,
			ExpiryTime: -1,
		}
	}
}

func NewValueWithExpiry(val interface{}, expiryTime int64) *Value {
	return &Value{
		Val:        val,
		ExpiryTime: expiryTime,
	}
}
