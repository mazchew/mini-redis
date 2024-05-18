package handlers

import (
	"github.com/codecrafters-io/redis-starter-go/app/protocol"
	"github.com/codecrafters-io/redis-starter-go/app/kvstore"
)

func HandleSet(kv *kvstore.KVStore, args []string) *protocol.RESPType {
	kv.Set(args[0], args[1])
	data := make([]string, 0)
	data = append(data, "OK")
	return &protocol.RESPType{DataType: protocol.SimpleString, Data: data}
}