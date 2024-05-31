package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/braveokafor/go-mail-api/internal/server"
	"github.com/braveokafor/go-mail-api/pkg/config"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("cannot load config: %s", err))
	}

	server := server.NewServer(*cfg)

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
