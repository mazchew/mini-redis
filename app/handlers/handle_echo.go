package handlers

import (
	"github.com/codecrafters-io/redis-starter-go/app/protocol"
)

func HandleEcho(args []string) *protocol.RESPType {
	data := make([]interface{}, len(args))
	for i, arg := range args {
		data[i] = arg
	}
	return &protocol.RESPType{DataType: protocol.Array, Data: data}
}