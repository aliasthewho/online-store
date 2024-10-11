package models

// ErrorResponse represents an error response
type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Details    string `json:"details"`
}

func NewErrorResponse(statusCode int, message string, details string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
		Details:    details,
	}
}
