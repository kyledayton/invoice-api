package web

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	*http.Server
	Handler http.Handler
	port    int
}

func NewServer(port int, handler http.Handler) *Server {
	addr := fmt.Sprintf(":%d", port)

	return &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
		port: port,
	}
}

func (s *Server) ListenAndServe() error {
	log.Printf("Invoice API is running on port %d", s.port)
	return s.Server.ListenAndServe()
}
