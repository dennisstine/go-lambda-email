package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dennisstine/go-lambda-email/pkg/responses"
	"github.com/dennisstine/go-lambda-email/pkg/ses"
	"github.com/dennisstine/go-lambda-email/pkg/structs"
	"github.com/dennisstine/go-lambda-email/pkg/validation"
	"net/http"
)

// Processes the incoming request and sends a response back to the client
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var message structs.Message

	// Parsing the request body into a Message instance
	unmarshalErr := json.Unmarshal([]byte(request.Body), &message)
	if unmarshalErr != nil {
		return responses.ErrorResponse(unmarshalErr, http.StatusInternalServerError)
	}

	validationErr := validation.Validate(message)
	if validationErr != nil {
		return responses.ErrorResponse(validationErr, http.StatusBadRequest)
	}

	sendError := ses.SendEmail(message)
	if sendError != nil {
		return responses.ErrorResponse(sendError, http.StatusInternalServerError)
	}

	return responses.SuccessResponse()
}

func main() {
	lambda.Start(handler)
}
