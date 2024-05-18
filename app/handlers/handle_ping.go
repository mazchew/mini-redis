package handlers

import (
	"github.com/codecrafters-io/redis-starter-go/app/protocol"
)


func HandlePing(args []string) *protocol.RESPType {
	data := make([]string, 0)
	data = append(data, "PONG")
	return &protocol.RESPType{DataType: protocol.SimpleString, Data: data}
}