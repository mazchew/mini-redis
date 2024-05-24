package handlers

import (
	"os"

	"github.com/codecrafters-io/redis-starter-go/app/config"
	"github.com/codecrafters-io/redis-starter-go/app/protocol"
)

func HandleKeys(config *config.Config, args []string) *protocol.RESPType {

	param := args[0]

	if param == "*" {
		fileContent := readFile(config.Dir + "/" + config.DbFilename)
		return &protocol.RESPType{
			DataType: protocol.Array,
			Data: []interface{}{
				&protocol.RESPType{DataType: protocol.BulkString, Data: []interface{}{fileContent}},
			},
		}
	}

	return &protocol.RESPType{
		DataType: protocol.Array,
		Data: []interface{}{
			&protocol.RESPType{DataType: protocol.BulkString, Data: []interface{}{nil}},
		},
	}

}

func readFile(path string) string {
	c, _ := os.ReadFile(path)
	key := parseTable(c)
	str := key[4 : 4+key[3]]
	return string(str)
}

func parseTable(bytes []byte) []byte {
	start := sliceIndex(bytes, protocol.OpCodeResizeDB)
	end := sliceIndex(bytes, protocol.OpCodeEOF)
	return bytes[start+1 : end]
}

func sliceIndex(data []byte, sep byte) int {
	for i, b := range data {
		if b == sep {
			return i
		}
	}

	return -1
}
