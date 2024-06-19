package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/braveokafor/go-mail-api/internal/types"
)

type Server struct {
	port   int
	config types.Config
}

func NewServer(cfg types.Config) *http.Server {
	newServer := &Server{
		port:   cfg.APIPort,
		config: cfg,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
