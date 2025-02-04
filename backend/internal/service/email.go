package service

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(toEmail *mail.Email, subject, body string) error {
	rootDir := filepath.Join("..", ".env")

	if err := godotenv.Load(rootDir); err != nil {
		log.Println("No .env file found, loading environment variables from system.")
	}
	apiKey := os.Getenv("SENDGRID_API_KEY")
	fromName := os.Getenv("FROM_NAME")
	fromEmail := os.Getenv("FROM_EMAIL")

	log.Printf("apiKey: %s; name: %s; from: %s", apiKey, fromName, fromEmail)

	if apiKey == "" || fromEmail == "" || fromName == "" {
		return fmt.Errorf("SendGrid API key or sender email not configured")
	}

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	from := mail.Email{Name: fromName, Address: fromEmail}
	message := mail.NewSingleEmail(&from, subject, toEmail, "", body)

	response, err := client.Send(message)

	if err != nil {
		log.Printf("SendGrid error: %v", err)
		return fmt.Errorf("Failed to send email: %v", err)
	}

	log.Printf("SendGrid response: %d - %s", response.StatusCode, response.Body)

	if response.StatusCode >= 400 {
		return fmt.Errorf("Failed to send email: status code %d, body: %s", response.StatusCode, response.Body)
	}

	return nil
}
