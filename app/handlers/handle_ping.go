package handlers

import (
	"fmt"
	"net"
)


func HandlePing(conn net.Conn) error {
	_, err := conn.Write([]byte("+PONG\r\n"))
	if err != nil {
		fmt.Println("Error writing:", err.Error())
		return err
	}
	return nil
}