package handlers

import (
	"fmt"
	
	"github.com/codecrafters-io/redis-starter-go/app/config"
	"github.com/codecrafters-io/redis-starter-go/app/protocol"
)

func HandleConfigGet(cfg *config.Config, args []string) *protocol.RESPType {
	
	fmt.Println("==== args: ", args)
	param := args[0]
	
	var value string
	switch param {
	case "dir":
		value = cfg.Dir
	case "dbfilename":
		value = cfg.DbFilename
	default:
		return &protocol.RESPType{
			DataType: protocol.BulkString,
			Data:     []interface{}{nil},
		}
	}

	return &protocol.RESPType{
		DataType: protocol.Array,
		Data: []interface{}{
			&protocol.RESPType{DataType: protocol.BulkString, Data: []interface{}{param}},
			&protocol.RESPType{DataType: protocol.BulkString, Data: []interface{}{value}},
		},
	}
}