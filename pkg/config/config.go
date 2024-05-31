package config

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port int
}

func LoadConfig() (*Config, error) {
	portStr, exists := os.LookupEnv("PORT")
	if !exists {
		portStr = "8080"
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("cannot convert port to int: %w", err)

	}

	return &Config{
		Port: port,
	}, nil
}
