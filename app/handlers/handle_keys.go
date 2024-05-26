package handlers

import (
	"github.com/codecrafters-io/redis-starter-go/app/kvstore"
	"github.com/codecrafters-io/redis-starter-go/app/protocol"
)

func HandleKeys(kv *kvstore.KVStore, args []string) *protocol.RESPType {

	param := args[0]

	var data []interface{}

	if param == "*" {
		for k := range kv.Store {
			data = append(data, &protocol.RESPType{DataType: protocol.BulkString, Data: []interface{}{k}})
		}

		return &protocol.RESPType{
			DataType: protocol.Array,
			Data:     data,
		}
	}

	return &protocol.RESPType{
		DataType: protocol.Array,
		Data: []interface{}{
			&protocol.RESPType{DataType: protocol.BulkString, Data: []interface{}{nil}},
		},
	}

}
