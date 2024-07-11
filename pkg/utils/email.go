package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/smtp"
	"os"
)

func SendEmail(to string, subject string, body string) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fromEmail := os.Getenv("FROM_EMAIL")
	appPassword := os.Getenv("EMAIL_APP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	toEmails := []string{to}
	msg := []byte("Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
		body + "\r\n")

	auth := smtp.PlainAuth("", fromEmail, appPassword, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, toEmails, msg)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
