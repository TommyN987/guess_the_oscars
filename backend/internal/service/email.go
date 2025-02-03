package service

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(toEmail *mail.Email, subject, body string) error {
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	fromName := os.Getenv("FROM_NAME")
	fromEmail := os.Getenv("FROM_EMAIL")
	from := mail.Email{Name: fromName, Address: fromEmail}
	message := mail.NewSingleEmail(&from, subject, toEmail, "", body)

	response, err := client.Send(message)

	if err != nil {
		return fmt.Errorf("Failed to send email: %v", err)
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("Failed to send email: status code %d, body: %s", response.StatusCode, response.Body)
	}

	return nil
}
