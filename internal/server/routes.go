package server

import (
	"net/http"

	"github.com/braveokafor/go-mail-api/internal/handlers"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HealthCheckHandler)
	mux.HandleFunc("/health", handlers.HealthCheckHandler)
	mux.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		handlers.EmailHandler(w, r, s.config)
	})

	return mux
}
