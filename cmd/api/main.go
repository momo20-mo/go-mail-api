package main

import (
	"fmt"
	"os"

	"log/slog"

	"github.com/braveokafor/go-mail-api/internal/config"
	"github.com/braveokafor/go-mail-api/internal/server"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	defer logger.Info("Server stopped")

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Cannot load config", slog.Any("error", err))
		os.Exit(1)
	}

	server := server.NewServer(*cfg)

	logger.Info(fmt.Sprintf("Listening on port %d", cfg.APIPort))

	if err := server.ListenAndServe(); err != nil {
		logger.Error("Cannot start server", slog.Any("error", err))
		os.Exit(1)
	}
}
