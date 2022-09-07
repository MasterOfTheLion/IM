package serv

import "sync"

type Server struct {
	once    sync.Once
	id      string
	address string
}