package todoback

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}
