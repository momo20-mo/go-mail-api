package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/braveokafor/go-mail-api/internal/email"

	"github.com/braveokafor/go-mail-api/internal/types"
)

func EmailHandler(w http.ResponseWriter, r *http.Request, cfg types.Config) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var emailRequest types.EmailRequest
	err := json.NewDecoder(r.Body).Decode(&emailRequest)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	mailService, err := email.New(cfg)
	if err != nil {
		http.Error(w, "Failed to initialize mail service", http.StatusInternalServerError)
		return
	}
	defer mailService.Close()

	err = mailService.SendEmail(emailRequest)
	if err != nil {
		log.Println(err)

		switch {
		case err.Error() == "only one of Text or HTML should be set":
			http.Error(w, "only one of Text or HTML should be set", http.StatusInternalServerError)
		default:
			http.Error(w, "Failed to send email", http.StatusInternalServerError)
		}

		//http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Email sent successfully",
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
