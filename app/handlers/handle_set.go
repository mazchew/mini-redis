package handlers

import (
	"strconv"
	"strings"
	"github.com/codecrafters-io/redis-starter-go/app/protocol"
	"github.com/codecrafters-io/redis-starter-go/app/kvstore"
)

func HandleSet(kv *kvstore.KVStore, args []string) *protocol.RESPType {
	var ttl int64 = -1
	if len(args) == 4 && strings.EqualFold(args[2], "px") {
		parsedTTL, _ := strconv.Atoi(args[3])
		ttl = int64(parsedTTL)
	}

	kv.Set(args[0], args[1], int64(ttl))
	data := make([]interface{}, 0)
	data = append(data, "OK")
	return &protocol.RESPType{DataType: protocol.SimpleString, Data: data}
}