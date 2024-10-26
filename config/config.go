package config

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type SMTPConfig struct {
	UseTLS   bool
	Port     int
	Host     string
	Username string
	Password string
}

type Config struct {
	MailAPIPort string
	SMTPConfig  SMTPConfig
}

func LoadConfig() *Config {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Define command-line flags with defaults from environment variables
	cfg := &Config{
		MailAPIPort: getEnv("MAIL_API_PORT", "8080"),
		SMTPConfig: SMTPConfig{
			Port:     getEnvAsInt("MAIL_SMTP_PORT", 25),
			Host:     getEnv("MAIL_SMTP_HOST", ""),
			Username: getEnv("MAIL_SMTP_USER", ""),
			Password: getEnv("MAIL_SMTP_PASS", ""),
			UseTLS:   getEnvAsBool("MAIL_SMTP_USE_TLS", false),
		},
	}

	flag.StringVar(&cfg.MailAPIPort, "api-port", cfg.MailAPIPort, "Port for the API server")
	flag.IntVar(&cfg.SMTPConfig.Port, "smtp-port", cfg.SMTPConfig.Port, "Port for the SMTP server")
	flag.StringVar(&cfg.SMTPConfig.Host, "smtp-host", cfg.SMTPConfig.Host, "Host for the SMTP server")
	flag.StringVar(&cfg.SMTPConfig.Username, "smtp-user", cfg.SMTPConfig.Username, "User for the SMTP server")
	flag.StringVar(&cfg.SMTPConfig.Password, "smtp-pass", cfg.SMTPConfig.Password, "Password for the SMTP server")
	flag.BoolVar(&cfg.SMTPConfig.UseTLS, "smtp-use-tls", cfg.SMTPConfig.UseTLS, "Use TLS for the SMTP server")

	flag.Parse()

	// Validate required fields
	if cfg.SMTPConfig.Host == "" || cfg.SMTPConfig.Username == "" || cfg.SMTPConfig.Password == "" {
		log.Fatal("SMTP host, user, and password are required.")
	}

	return cfg
}

// getEnv retrieves an environment variable or returns a default value if not set
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt retrieves an environment variable as an integer or returns a default value if not set
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		log.Printf("Invalid integer value for %s, using default: %d", key, defaultValue)
	}
	return defaultValue
}

// getEnvAsBool retrieves an environment variable as a boolean or returns a default value if not set
func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		return (value == "true")
	}
	return defaultValue
}
