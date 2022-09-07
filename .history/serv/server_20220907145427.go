package serv

import (
	"net"
	"sync"
)

type Server struct {
	once    sync.Once
	id      string
	address string
	sync.Mutex
	users map[string]net.Conn
}

func NewServer(id, address string) *Server {
	return newServer(id, address)
}

func newServer(id, address string) *Server {
	return &Server{
		id: id,
		address: address,
		users: make(map[string]net.Conn, 100),
	}
}