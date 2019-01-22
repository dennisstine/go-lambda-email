package structs

// Representation of an incoming message payload
type Message struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
}
