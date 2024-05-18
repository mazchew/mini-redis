package main

import (
	"fmt"
	"net"
	"os"
	
	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/kvstore"
)

func handleConnection(kvStore *kvstore.KVStore, conn net.Conn) {
	defer conn.Close()

	parser := resp.NewParser(conn)
	encoder := resp.NewEncoder(conn)

	for {
		command, _ := parser.ParseCommand()
		if command != nil {
			output := resp.ExecuteCommand(kvStore, command)
			err := encoder.Write(output)
			if err != nil {
				fmt.Println("Error writing to client")
			}
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	fmt.Println("Server started...")
	kvStore := kvstore.NewKVStore()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(kvStore, conn)
	}
}