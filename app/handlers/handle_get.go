package handlers

import (
	"github.com/codecrafters-io/redis-starter-go/app/protocol"
	"github.com/codecrafters-io/redis-starter-go/app/kvstore"
)

func HandleGet(kv *kvstore.KVStore, args []string) *protocol.RESPType {
	data := make([]string, 0)
	val, err := kv.Get(args[0])
	if err != nil {
		data = append(data, "-1")
	} else {
		data = append(data, val)
	}

	return &protocol.RESPType{DataType: protocol.BulkString, Data: data}
}