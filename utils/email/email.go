package email

import (
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendOne(subject string, toName string, toEmail string, plainTextContent string, htmlContent string) (bool, error) {
	// Making mail content
	from := mail.NewEmail("Badreddin", "badreddin.laabed@gmail.com")
	to := mail.NewEmail(toName, toEmail)
	// Sending
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
