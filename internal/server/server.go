package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/braveokafor/go-mail-api/pkg/config"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
}

func NewServer(cfg config.Config) *http.Server {
	NewServer := &Server{
		port: cfg.Port,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
