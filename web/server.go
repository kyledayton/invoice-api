package web

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	*http.Server
	port int
}

func NewServer(port int) *Server {
	routes := makeRoutes()
	addr := fmt.Sprintf(":%d", port)

	return &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: routes,
		},
		port: port,
	}
}

func (s *Server) ListenAndServe() error {
	log.Printf("Invoice API is running on port %d", s.port)
	return s.Server.ListenAndServe()
}
