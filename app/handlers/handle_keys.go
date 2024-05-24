package handlers

import (
	"github.com/codecrafters-io/redis-starter-go/app/config"
	"github.com/codecrafters-io/redis-starter-go/app/protocol"
	"github.com/codecrafters-io/redis-starter-go/app/utils"
)

func HandleKeys(config *config.Config, args []string) *protocol.RESPType {

	param := args[0]

	if param == "*" {
		fileContent := utils.ParseFile(config.Dir + "/" + config.DbFilename)

		var data []interface{}

		for _, kv := range fileContent {
			data = append(data, &protocol.RESPType{DataType: protocol.BulkString, Data: []interface{}{kv.Key}})
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
