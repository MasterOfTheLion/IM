package serv

import (
	"net"
	"net/http"
	"sync"

	"github.com/gobwas/ws"
	"github.com/sirupsen/logrus"
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

func (s *Server) Start() error {
	mux := http.NewServeMux()
	log := logrus.WithFields(logrus.Fields{
		"module": "Server",
		"listen": s.address,
		"id": s.id,
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			conn.Close()
			return
		}
		user := r.URL.Query().Get("user")
		if user == "" {
			conn.Close()
			return
		}
		old, ok := s.addUser(user, conn)
		if ok {
			old.Close()
		}
		log.Infof("user %s in ", user)

		go func(user string, conn net.Conn) {
			err := s.readloop(user, conn)
			if err != nil {
				log.Error(err)
			}
			conn.Close()
			s.delUser(user)
			log.Infof("connection of %s cloesd ", user)
		}(user, conn)
	})
	log.Infof("started")
	return http.ListenAndServe(s.address, mux)
}