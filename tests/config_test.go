package tests

import (
	"os"
	"testing"

	"github.com/braveokafor/go-mail-api/config"
)

func TestConfig(t *testing.T) {
	os.Setenv("MAIL_API_PORT", "8081")
	os.Setenv("MAIL_SMTP_PORT", "587")
	os.Setenv("MAIL_SMTP_HOST", "smtp.gmail.com")
	os.Setenv("MAIL_SMTP_USER", "example@gmail.com")
	os.Setenv("MAIL_SMTP_PASS", "12345")
	os.Setenv("MAIL_SMTP_USE_TLS", "true")

	defer os.Clearenv()

	cfg := config.LoadConfig()

	if cfg.MailAPIPort != "8081" {
		t.Errorf("expected port to be 8081; got %v", cfg.MailAPIPort)
	}

	if cfg.SMTPConfig.Port != 587 {
		t.Errorf("expected port to be 587; got %v", cfg.SMTPConfig.Port)
	}

	if cfg.SMTPConfig.Host != "smtp.gmail.com" {
		t.Errorf("expected host to be smtp.gmail.com; got %v", cfg.SMTPConfig.Host)
	}

	if cfg.SMTPConfig.Username != "example@gmail.com" {
		t.Errorf("expected user to be example@gmail.com ; got %v", cfg.SMTPConfig.Username)
	}

	if cfg.SMTPConfig.Password != "12345" {
		t.Errorf("expected password to be 12345; got %v", cfg.SMTPConfig.Password)
	}

	if cfg.SMTPConfig.UseTLS != true {
		t.Errorf("expected useTLS to be true; got %v", cfg.SMTPConfig.UseTLS)
	}
}
