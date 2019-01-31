package responses

import (
	"net/http"
)

var header = map[string]string{"Content-Type": "application/json"}

// Return a success (200) to the client with nil error
func SuccessResponse() (LambdaResponse, error) {

	return LambdaResponse{
		Headers:    header,
		StatusCode: http.StatusOK,
	}, nil
}

// Return an error and status code (e.g. 400, 500) to the client
func ErrorResponse(code int, err error) (LambdaResponse, error) {

	return LambdaResponse{
		Headers:    header,
		StatusCode: code,
		Body:       http.StatusText(code),
	}, err
}
