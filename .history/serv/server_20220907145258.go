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