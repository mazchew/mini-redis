package handlers

import (
	"fmt"
	"net"
)

func HandleEcho(conn net.Conn, input string) error {
	_, err := conn.Write([]byte("+" + input + "\r\n"))
	if err != nil {
		fmt.Println("Error writing:", err.Error())
		return err
	}

	return nil
}