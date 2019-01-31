package responses

// Representation of a response to the client
type LambdaResponse struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}
