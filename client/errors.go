package client

import "fmt"

// APIError represents an error returned by the ToyyibPay API.
type APIError struct {
	StatusCode int
	Message    string
}

// Error returns the error message.
func (e *APIError) Error() string {
	return fmt.Sprintf("API Error: %s (Status Code: %d)", e.Message, e.StatusCode)
}
