package utils

import (
	"github.com/joho/godotenv"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"os"
)

func SendEmail(to, subject, body string, attachments []string) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	e := email.NewEmail()
	e.From = os.Getenv("EMAIL_FROM")
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(body)

	for _, attachmentPath := range attachments {
		_, err := e.AttachFile(attachmentPath)
		if err != nil {
			return err
		}
	}

	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	err = e.Send(smtpServer+":"+smtpPort, smtp.PlainAuth("", smtpUser, smtpPassword, smtpServer))
	return err
}