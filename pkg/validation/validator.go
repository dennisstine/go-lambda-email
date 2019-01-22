package validation

import (
	"errors"
	"github.com/dennisstine/go-lambda-email/pkg/structs"
)

// Client error responses
var (
	nameNotProvided    = errors.New("name not provided")
	emailNotProvided   = errors.New("email not provided")
	subjectNotProvided = errors.New("subject not provided")
	textNotProvided    = errors.New("text not provided")
)

// Return an error message if any fields are empty; nil otherwise
func Validate(message structs.Message) error {

	if len(message.Name) < 1 {
		return nameNotProvided
	} else if len(message.Email) < 1 {
		return emailNotProvided
	} else if len(message.Subject) < 1 {
		return subjectNotProvided
	} else if len(message.Text) < 1 {
		return textNotProvided
	} else {
		return nil
	}
}
