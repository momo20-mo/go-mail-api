package tests

import (
	"os"
	"testing"

	"github.com/braveokafor/go-mail-api/internal/config"
)

func TestConfig(t *testing.T) {
	os.Setenv("MAIL_API_PORT", "8081")
	os.Setenv("MAIL_SMTP_PORT", "587")
	os.Setenv("MAIL_SMTP_HOST", "smtp.gmail.com")
	os.Setenv("MAIL_SMTP_USER", "example@gmail.com")
	os.Setenv("MAIL_SMTP_PASS", "12345")
	os.Setenv("MAIL_SMTP_USE_TLS", "True")

	defer os.Clearenv()

	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("error loading config: %v", err)
	}

	if cfg.APIPort != 8081 {
		t.Errorf("expected port to be 8081; got %v", cfg.APIPort)
	}

	if cfg.SMTPPort != 587 {
		t.Errorf("expected port to be 587; got %v", cfg.SMTPPort)
	}

	if cfg.SMTPHost != "smtp.gmail.com" {
		t.Errorf("expected host to be smtp.gmail.com; got %v", cfg.SMTPHost)
	}

	if cfg.SMTPUser != "example@gmail.com" {
		t.Errorf("expected user to be example@gmail.com ; got %v", cfg.SMTPUser)
	}

	if cfg.SMTPPass != "12345" {
		t.Errorf("expected password to be 12345; got %v", cfg.SMTPPass)
	}

	if cfg.SMTPUseTLS != true {
		t.Errorf("expected useTLS to be true; got %v", cfg.SMTPUseTLS)
	}
}
