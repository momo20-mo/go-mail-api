package main

import (
	"fmt"
	"net/http"

	"github.com/braveokafor/go-mail-api/config"
	"github.com/braveokafor/go-mail-api/handlers"
)

func main() {
	cfg := config.LoadConfig()

	handlers.SetConfig(cfg)

	http.HandleFunc("/", handlers.HealthCheck)
	http.HandleFunc("/healthz", handlers.HealthCheck)
	http.HandleFunc("/send", handlers.SendEmail)

	fmt.Printf("Starting server on :%s\n", cfg.MailAPIPort)
	if err := http.ListenAndServe(":"+cfg.MailAPIPort, nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
