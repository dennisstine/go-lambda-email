package responses

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

var header = map[string]string{"Content-Type": "application/json"}

// Return a success (200) to the client
func SuccessResponse() (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		Headers:    header,
		StatusCode: http.StatusOK,
	}, nil
}

// Return an error (e.g. 400, 500) to the client with the error message
func ErrorResponse(err error, code int) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		Headers:    header,
		StatusCode: code,
		Body:       err.Error(),
	}, nil
}
