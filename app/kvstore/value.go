package kvstore

import "time"

type Value struct {
	Val interface{}
	Ttl int64
	SetAt int64
}

func NewValue(val interface{}, ttl int64) *Value {
	return &Value{
		Val: val,
		Ttl: ttl,
		SetAt: time.Now().UnixNano() / 1e6,
	}
}
