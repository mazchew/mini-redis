package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/codecrafters-io/redis-starter-go/app/config"
	"github.com/codecrafters-io/redis-starter-go/app/kvstore"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

type Server struct {
	kvStore  *kvstore.KVStore
	listener net.Listener
	config   *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		kvStore: kvstore.NewKVStore(),
		config:  cfg,
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	parser := resp.NewParser(conn)
	encoder := resp.NewEncoder(conn)

	for {
		command, _ := parser.ParseCommand()
		if command != nil {
			output := resp.ExecuteCommand(s.kvStore, s.config, command)
			err := encoder.Write(output)
			if err != nil {
				fmt.Println("Error writing to client")
			}
		}
	}
}

func (s *Server) Start() error {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		return fmt.Errorf("failed to bind to port 6379: %v", err)
	}
	s.listener = l

	log.Println("Server started...")

	go s.acceptConnections()
	return nil
}

func (s *Server) acceptConnections() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			// Check if the error is because the listener was closed.
			if netErr, ok := err.(net.Error); ok && !netErr.Temporary() {
				log.Println("Listener closed. Stopping acceptance of new connections.")
				break
			}
			log.Println("Error accepting connection:", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) Stop() {
	if s.listener != nil {
		s.listener.Close()
	}
}

func main() {
	dir := flag.String("dir", "/Users/maz/Desktop/Code/tmp/redis-files", "directory to store RDB files")
	dbFilename := flag.String("dbfilename", "dump.rdb", "filename for the RDB file")
	flag.Parse()

	cfg := config.NewConfig(*dir, *dbFilename)

	server := NewServer(cfg)

	if err := server.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down server...")
	server.Stop()
	log.Println("Server stopped gracefully")
}
