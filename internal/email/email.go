package email

import (
	"errors"
	"fmt"
	"log"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"

	"github.com/braveokafor/go-mail-api/internal/types"
)

type Service interface {
	SendEmail(types.EmailRequest) error

	Close() error
}

type service struct {
	client *mail.SMTPClient
}

func New(cfg types.Config) (Service, error) {
	server := mail.NewSMTPClient()
	server.Host = cfg.SMTPHost
	server.Port = cfg.SMTPPort
	server.Username = cfg.SMTPUser
	server.Password = cfg.SMTPPass
	server.Encryption = mail.EncryptionTLS
	server.KeepAlive = true
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	if !cfg.SMTPUseTLS {
		server.Encryption = mail.EncryptionNone
	}

	smtpClient, err := server.Connect()
	if err != nil {
		log.Println("failed to connect to mail server")

		return nil, fmt.Errorf("failed to connect to mail server: %w", err)
	}

	return &service{client: smtpClient}, nil
}

func (s *service) SendEmail(emailRequest types.EmailRequest) error {
	// only one of Text or HTML should be set.
	bodyFieldsSet := 0
	if emailRequest.Text != "" {
		bodyFieldsSet++
	}
	if emailRequest.HTML != "" {
		bodyFieldsSet++
	}

	//  Return an error if more than one is set
	if bodyFieldsSet > 1 {
		return errors.New("only one of Text or HTML should be set")
	}

	email := mail.NewMSG()
	email.SetFrom(emailRequest.From)
	email.AddTo(emailRequest.To...)
	email.AddCc(emailRequest.CC...)
	email.AddBcc(emailRequest.BCC...)
	email.SetSubject(emailRequest.Subject)

	// Set custom headers
	for key, value := range emailRequest.Headers {
		email.AddHeader(key, value)
	}

	// Set email body
	if emailRequest.Text != "" {
		email.SetBody(mail.TextPlain, emailRequest.Text)
	} else if emailRequest.HTML != "" {
		email.SetBody(mail.TextHTML, emailRequest.HTML)
	}

	// Set email priority
	switch emailRequest.Priority {
	case "High":
		email.SetPriority(mail.PriorityHigh)
	case "Low":
		email.SetPriority(mail.PriorityLow)
	}

	// Add attachments
	for _, attachment := range emailRequest.Attachments {
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

	err := email.Send(s.client)
	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	log.Printf("Email sent: %s \n", emailRequest.Subject)
	return nil

}

func (s *service) Close() error {
	if s.client == nil {
		return nil
	}

	err := s.client.Close()
	if err != nil {
		return fmt.Errorf("error closing mail client: %w", err)
	}

	s.client = nil
	log.Println("Disconnected from mail client")
	return nil
}
