package ses

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/dennisstine/go-lambda-email/pkg/structs"
	"os"
)

var toEmail string
var fromEmail string
var emailClient *ses.SES

// Retrieve ENV variables from AWS config and establish the SES client
func init() {

	toEmail = os.Getenv("TO_EMAIL")
	fromEmail = os.Getenv("FROM_EMAIL")

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_DATACENTER"))},
	)

	emailClient = ses.New(sess)
}

// Send the email via SES
func SendEmail(message structs.Message) error {

	emailParams := &ses.SendEmailInput{
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String("From: " + message.Name + " - " + message.Email + "\n\n" + message.Text),
				},
			},
			Subject: &ses.Content{
				Data: aws.String("Website Form Message - " + message.Subject),
			},
		},
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(toEmail)},
		},
		Source: aws.String(fromEmail),
	}

	_, err := emailClient.SendEmail(emailParams)

	return err
}
