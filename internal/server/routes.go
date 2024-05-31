package server

import (
	"net/http"

	"github.com/braveokafor/go-mail-api/internal/handlers"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HelloWorldHandler)
	mux.HandleFunc("/health", handlers.HealthCheckHandler)

	return mux
}
