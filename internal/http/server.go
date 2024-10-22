package http

import (
	"net/http"
)

type Server struct {
	http.Server
}

func NewServer(addr string) *Server {
	httpServer := &Server{}
	httpServer.Server.Addr = addr

	return httpServer
}

func (srv *Server) WithHandler(h http.Handler) *Server {
	srv.Server.Handler = h
	return srv
}
