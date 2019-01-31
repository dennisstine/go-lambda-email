package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dennisstine/go-lambda-email/pkg/responses"
	"github.com/dennisstine/go-lambda-email/pkg/ses"
	"github.com/dennisstine/go-lambda-email/pkg/structs"
	"github.com/dennisstine/go-lambda-email/pkg/validation"
	"net/http"
)

// On message, processes the request and returns the appropriate response
func handler(message structs.Message) (responses.LambdaResponse, error) {

	validationErr := validation.Validate(message)
	if validationErr != nil {
		return responses.ErrorResponse(http.StatusBadRequest, validationErr)
	}

	sesErr := ses.SendEmail(message)
	if sesErr != nil {
		return responses.ErrorResponse(http.StatusInternalServerError, sesErr)
	}

	return responses.SuccessResponse()
}

func main() {
	lambda.Start(handler)
}
