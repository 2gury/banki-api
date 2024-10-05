package server

import (
	"net/http"
	"time"
)

type HTTPServer struct {
	server *http.Server
}

func NewHTTPServer(handler http.Handler, port string) *HTTPServer {
	return &HTTPServer{
		server: &http.Server{
			Addr:         ":" + port,
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		},
	}
}

func (s *HTTPServer) Launch() error {
	return s.server.ListenAndServe()
}
