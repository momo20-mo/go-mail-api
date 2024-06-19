package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/braveokafor/go-mail-api/internal/types"
	_ "github.com/joho/godotenv/autoload"
)

func LoadConfig() (*types.Config, error) {
	// Set default values
	apiPort := 8080
	smtpPort := 25
	smtpUseTLS := false

	// Parse environment variables
	if val := os.Getenv("MAIL_API_PORT"); val != "" {
		port, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("invalid value for MAIL_API_PORT: %w", err)
		}
		apiPort = port
	}

	if val := os.Getenv("MAIL_SMTP_PORT"); val != "" {
		port, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("invalid value for MAIL_SMTP_PORT: %w", err)
		}
		smtpPort = port
	}

	if val := os.Getenv("MAIL_SMTP_USE_TLS"); val != "" {
		useTLS, err := strconv.ParseBool(val)
		if err != nil {
			return nil, fmt.Errorf("invalid value for MAIL_SMTP_USE_TLS: %w", err)
		}
		smtpUseTLS = useTLS
	}

	smtpHost := os.Getenv("MAIL_SMTP_HOST")
	smtpUser := os.Getenv("MAIL_SMTP_USER")
	smtpPass := os.Getenv("MAIL_SMTP_PASS")

	// Parse command line flags
	flag.IntVar(&apiPort, "api-port", apiPort, "Port for the API server")
	flag.IntVar(&smtpPort, "smtp-port", smtpPort, "Port for the SMTP server")
	flag.StringVar(&smtpHost, "smtp-host", smtpHost, "Host for the SMTP server")
	flag.StringVar(&smtpUser, "smtp-user", smtpUser, "User for the SMTP server")
	flag.StringVar(&smtpPass, "smtp-pass", smtpPass, "Password for the SMTP server")
	flag.BoolVar(&smtpUseTLS, "smtp-use-tls", smtpUseTLS, "Use TLS for the SMTP server")
	flag.Parse()

	// Check if all required configuration values are set
	if smtpHost == "" || smtpUser == "" || smtpPass == "" {
		return nil, fmt.Errorf("missing required configuration values")
	}

	return &types.Config{
		APIPort:    apiPort,
		SMTPPort:   smtpPort,
		SMTPHost:   smtpHost,
		SMTPUser:   smtpUser,
		SMTPPass:   smtpPass,
		SMTPUseTLS: smtpUseTLS,
	}, nil
}
