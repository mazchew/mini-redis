package resp

import (
	"fmt"
	
	"github.com/codecrafters-io/redis-starter-go/app/protocol"
	"github.com/codecrafters-io/redis-starter-go/app/handlers"
	"github.com/codecrafters-io/redis-starter-go/app/kvstore"
	"github.com/codecrafters-io/redis-starter-go/app/config"
)

type CommandName string

const (
	ECHO CommandName = "ECHO"
	PING CommandName = "PING"
	SET  CommandName = "SET"
	GET  CommandName = "GET"
	CONFIG_GET CommandName = "CONFIG GET"
)

type Command struct {
	Name CommandName
	Args []string
}

func ExecuteCommand(kv *kvstore.KVStore, cfg *config.Config, command *Command) *protocol.RESPType {
	if command == nil {
		return nil
	}

	fmt.Println("==== command: ", command)

	switch command.Name {
	case PING:
		return handlers.HandlePing(command.Args)
	case ECHO:
		return handlers.HandleEcho(command.Args)
	case SET:
		return handlers.HandleSet(kv, command.Args)
	case GET:
		return handlers.HandleGet(kv, command.Args)
	case CONFIG_GET:
		return handlers.HandleConfigGet(cfg, command.Args)
	}
	return nil
}