package resp

import (
	"github.com/codecrafters-io/redis-starter-go/app/config"
	"github.com/codecrafters-io/redis-starter-go/app/handlers"
	"github.com/codecrafters-io/redis-starter-go/app/kvstore"
	"github.com/codecrafters-io/redis-starter-go/app/protocol"
)

type CommandName string

const (
	ECHO       CommandName = "ECHO"
	PING       CommandName = "PING"
	SET        CommandName = "SET"
	GET        CommandName = "GET"
	CONFIG_GET CommandName = "CONFIG GET"
	KEYS       CommandName = "KEYS"
)

type Command struct {
	Name CommandName
	Args []string
}

func ExecuteCommand(kv *kvstore.KVStore, cfg *config.Config, command *Command) *protocol.RESPType {
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
	case CONFIG_GET:
		return handlers.HandleConfigGet(cfg, command.Args)
	case KEYS:
		return handlers.HandleKeys(kv, command.Args)
	}
	return nil
}
