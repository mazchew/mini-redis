package handlers

import (
	"github.com/codecrafters-io/redis-starter-go/app/protocol"
	"github.com/codecrafters-io/redis-starter-go/app/kvstore"
)

func HandleGet(kv *kvstore.KVStore, args []string) *protocol.RESPType {
	data := make([]interface{}, 0)
	val, err := kv.Get(args[0])
	if err != nil {
		return &protocol.RESPType{DataType: protocol.BulkString, Data: []interface{}{"-1"}}
	} else {
		data = append(data, val.Val)
	}

	return &protocol.RESPType{DataType: protocol.BulkString, Data: data}
}