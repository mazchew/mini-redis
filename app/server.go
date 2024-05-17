// package main
  
// import (
// 	"fmt"
// 	"net"
// 	"os"
// )

// func main() {
// 	l, err := net.Listen("tcp", "0.0.0.0:6379")
// 	if err != nil {
// 		fmt.Println("Failed to bind to port 6379")
// 		os.Exit(1)
// 	}
// 	defer l.Close()

// 	errChan := make(chan error, 0)
// 	for {
// 		go func() {
// 			conn, err := l.Accept()
// 			if err != nil {
// 				fmt.Println("Error accepting connection: ", err.Error())
// 				errChan <- err
// 			}
// 			defer conn.Close()

// 			fmt.Println("Connection established")
// 			for {
// 				err := handleConnection(conn)
// 				if err != nil {
// 					fmt.Println("Error handling connection: ", err.Error())
// 					errChan <- err
// 				}
// 			}
// 		}()
// 		select {
// 		case err := <-errChan:
// 			fmt.Println("Error: ", err.Error())
// 		default:
// 		}
// 	}
// }
  
// func handleConnection(conn net.Conn) error {
// 	bufIn := make([]byte, 1024)
// 	reqLen, err := conn.Read(bufIn)
// 	if err != nil {
// 		fmt.Println("Error reading:", err.Error())
// 		return err
// 	}

// 	resp, err := NewRESP(bufIn[:reqLen])
// 	if err != nil {
// 		fmt.Println("Error parsing:", err.Error())
// 		return err
// 	}
// 	fmt.Println("Received: ", resp.Command, " ", resp.Input)

// 	switch resp.Command {
// 	case "PING":
// 		err = handlePing(conn)
// 		if err != nil {
// 			fmt.Println("Error handling PING:", err.Error())
// 			return err
// 		}
// 	case "ECHO":
// 		err = handleEcho(conn, resp.Input)
// 		if err != nil {
// 			fmt.Println("Error handling ECHO:", err.Error())
// 			return err
// 		}
// 	default:
// 		fmt.Println("Unknown command: ", resp.Command)
// 	}

// 	return nil
// }

// func handleEcho(conn net.Conn, input string) error {
// 	_, err := conn.Write([]byte("+" + input + "\r\n"))
// 	if err != nil {
// 		fmt.Println("Error writing:", err.Error())
// 		return err
// 	}

// 	return nil
// }

// func handlePing(conn net.Conn) error {
// 	_, err := conn.Write([]byte("+PONG\r\n"))
// 	if err != nil {
// 		fmt.Println("Error writing:", err.Error())
// 		return err
// 	}
// 	return nil
// }

package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    l, err := net.Listen("tcp", "0.0.0.0:6379")
    if err != nil {
        fmt.Println("Failed to bind to port 6379")
        os.Exit(1)
    }
    defer l.Close()

    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting connection: ", err.Error())
            continue
        }

        go func(conn net.Conn) {
            defer conn.Close()  // Ensure connection is closed when the goroutine finishes

            fmt.Println("Connection established")
            for {
                if err := handleConnection(conn); err != nil {
                    fmt.Println("Error handling connection: ", err.Error())
                    return  // Exit the goroutine when an error occurs
                }
            }
        }(conn)
    }
}

func handleConnection(conn net.Conn) error {
    bufIn := make([]byte, 1024)
    reqLen, err := conn.Read(bufIn)
    if err != nil {
        fmt.Println("Error reading:", err.Error())
        return err
    }

    resp, err := NewRESP(bufIn[:reqLen])
    if err != nil {
        fmt.Println("Error parsing:", err.Error())
        return err
    }
    fmt.Println("Received: ", resp.Command, " ", resp.Input)

    switch resp.Command {
    case "PING":
        err = handlePing(conn)
        if err != nil {
            fmt.Println("Error handling PING:", err.Error())
            return err
        }
    case "ECHO":
        err = handleEcho(conn, resp.Input)
        if err != nil {
            fmt.Println("Error handling ECHO:", err.Error())
            return err
        }
    default:
        fmt.Println("Unknown command: ", resp.Command)
    }

    return nil
}

func handleEcho(conn net.Conn, input string) error {
    _, err := conn.Write([]byte("+" + input + "\r\n"))
    if err != nil {
        fmt.Println("Error writing:", err.Error())
        return err
    }
    return nil
}

func handlePing(conn net.Conn) error {
    _, err := conn.Write([]byte("+PONG\r\n"))
    if err != nil {
        fmt.Println("Error writing:", err.Error())
        return err
    }
    return nil
}
