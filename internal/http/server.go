package http

import (
	"net/http"
)

func NewServer(addr string) *http.Server {
	httpServer := &http.Server{Addr: addr}

	return httpServer
}
