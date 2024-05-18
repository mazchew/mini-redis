package handlers

import (
	"github.com/codecrafters-io/redis-starter-go/app/protocol"
)

func HandleEcho(args []string) *protocol.RESPType {
	data := make([]string, 0)
	data = append(data, args[0])
	return &protocol.RESPType{DataType: protocol.BulkString, Data: data}
}