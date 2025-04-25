package dto

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return e.Message
}

type ErrorResponse struct {
	Error string `json:"error"`
}
