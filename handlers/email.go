package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/braveokafor/go-mail-api/config"
	"github.com/braveokafor/go-mail-api/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

var (
	smtpClient *mail.SMTPClient
	once       sync.Once
	cfg        *config.Config
)

func InitializeSMTPClient() {
	smtpServer := mail.NewSMTPClient()

	smtpServer.Host = cfg.SMTPConfig.Host
	smtpServer.Port = cfg.SMTPConfig.Port
	smtpServer.Username = cfg.SMTPConfig.Username
	smtpServer.Password = cfg.SMTPConfig.Password
	smtpServer.KeepAlive = true
	smtpServer.ConnectTimeout = 30 * time.Second
	smtpServer.SendTimeout = 30 * time.Second
	smtpServer.Encryption = mail.EncryptionNone
	if cfg.SMTPConfig.UseTLS {
		smtpServer.Encryption = mail.EncryptionTLS
	}

	var err error
	log.Println("Connecting to SMTP server...")
	smtpClient, err = smtpServer.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to SMTP server: %v", err)
	}
}

func SendEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Invalid request method"}`, http.StatusMethodNotAllowed)
		return
	}

	var emailReq models.EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&emailReq); err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	once.Do(InitializeSMTPClient)

	email := mail.NewMSG()
	email.SetFrom(emailReq.From).
		AddTo(emailReq.To...).
		AddCc(emailReq.CC...).
		AddBcc(emailReq.BCC...).
		SetSubject(emailReq.Subject)

	switch emailReq.Priority {
	case "high":
		email.SetPriority(mail.PriorityHigh)
	case "low":
		email.SetPriority(mail.PriorityLow)
	}

	for key, value := range emailReq.Headers {
		email.AddHeader(key, value)
	}

	if emailReq.Text != "" {
		email.SetBody(mail.TextPlain, emailReq.Text)
	} else if emailReq.HTML != "" {
		email.SetBody(mail.TextHTML, emailReq.HTML)
	}

	for _, attachment := range emailReq.Attachments {
		file := &mail.File{
			Name:     attachment.Filename,
			MimeType: attachment.ContentType,
			Inline:   false,
		}
		if attachment.Encoded {
			file.B64Data = attachment.Content
		} else {
			file.Data = []byte(attachment.Content)
		}
		email.Attach(file)
	}

	log.Println("Sending email to:", emailReq.To)
	if err := email.Send(smtpClient); err != nil {
		log.Println("Failed to send email to", emailReq.To, ":", err)
		http.Error(w, `{"error": "Failed to send email"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Email sent successfully"}`))
}
