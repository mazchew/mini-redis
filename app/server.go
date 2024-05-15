package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

const (
	pongResponse  = "+PONG\r\n"
	unknownCmdMsg = "-ERR unknown command\r\n"
)

func main() {
	fmt.Println("Server is starting...")

	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Printf("Failed to bind to port 6379: %s\n", err)
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %s\n", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		command, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Read error: %s\n", err)
			continue
		}
		response := processCommand(command)
		if response != "" {
			_, err := conn.Write([]byte(response))
			if err != nil {
				fmt.Printf("Write error: %s\n", err)
				return
			}
		}
	}
}

func processCommand(command string) string {
	parts := strings.Fields(strings.TrimSpace(command))
	if len(parts) < 1 {
		return unknownCmdMsg
	}

	cmd := strings.ToUpper(parts[0])
	switch cmd {
	case "PING":
		return pongResponse
	case "ECHO":
		if len(parts) > 1 {
			return fmt.Sprintf("$%d\r\n%s\r\n", len(parts[1]), parts[1])
		}
		return "$-1\r\n"
	default:
		return unknownCmdMsg
	}
}