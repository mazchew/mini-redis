package resp

import (
	// "fmt"
	"github.com/codecrafters-io/redis-starter-go/app/protocol"
	"github.com/codecrafters-io/redis-starter-go/app/handlers"
	"github.com/codecrafters-io/redis-starter-go/app/kvstore"
)

type CommandName string

const (
	ECHO CommandName = "ECHO"
	PING CommandName = "PING"
	SET  CommandName = "SET"
	GET  CommandName = "GET"
)

type Command struct {
	Name CommandName
	Args []string
}

func ExecuteCommand(kv *kvstore.KVStore, command *Command) *protocol.RESPType {
	if command == nil {
		return nil
	}

	switch command.Name {
	case PING:
		return handlers.HandlePing(command.Args)
	case ECHO:
		return handlers.HandleEcho(command.Args)
	case SET:
		return handlers.HandleSet(kv, command.Args)
	case GET:
		return handlers.HandleGet(kv, command.Args)

	}
	return nil
}

/*

		// data := make([]string, 0)
		// val, err := s.Get(command.Args[0])
		// fmt.Println("command arg", command.Args[0])
		// if err != nil {
		// 	data = append(data, "-1")
		// } else {
		// 	data = append(data, val)
		// }
		// return &RESPType{DataType: BulkString, Data: data}

				// s.Set(command.Args[0], command.Args[1])
		// data := make([]string, 0)
		// data = append(data, "OK")
		// return &RESPType{DataType: SimpleString, Data: data}

*/
