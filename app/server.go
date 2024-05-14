package main

import (
	"fmt"
	"io"
	"net"
	"os"
 	"strings"
)


func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		if strings.Contains(string(buffer[:n]), "PING") {
			_, err := conn.Write([]byte("+PONG\r\n"))
			if err != nil {
				return
			}

		}
	}
}

