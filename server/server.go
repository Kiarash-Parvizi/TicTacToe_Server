package server

import (
	"fmt"
	"net"
	"strconv"
)

func init() {
}

// ListenAndServe : creates a new server and starts listening for incomming connections on the provided port
func ListenAndServe(port int) error {
	newServer(port).listen()
	return nil
}

// Server : game-server
type Server struct {
	port     int
	listener net.Listener
}

// NewServer : creates a new server
func newServer(port int) Server {
	server := Server{port: port}
	l, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	println("newServer:", l.Addr().String())
	server.listener = l
	return server
}

// Listen : listen for connections on this server
func (s Server) listen() {
	for {
		fmt.Println("waiting for connection...")
		connection, err := s.listener.Accept()
		fmt.Println("new connection\n\t")
		if err != nil {
			continue
		}
		go handleConnection(connection)
	}
}
