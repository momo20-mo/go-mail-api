package tests

import (
	"os"
	"testing"

	"github.com/braveokafor/go-mail-api/pkg/config"
)

func TestConfig(t *testing.T) {
	os.Setenv("PORT", "8080")
	defer os.Clearenv()

	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("error loading config: %v", err)
	}

	if cfg.Port != 8080 {
		t.Errorf("expected port to be 8080; got %v", cfg.Port)
	}
}
