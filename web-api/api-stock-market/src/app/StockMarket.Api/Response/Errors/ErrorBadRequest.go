package Errors

// ErrorBadRequest represents the 400 error.
// swagger:model
type ErrorBadRequest struct {
	// Error code.
	// example: 400
	Code int `json:"code" example:"400"`
	// Error description.
	// example: "Bad request"
	Error string `json:"error" example:"Bad request"`
	// Error message.
	// example: "Invalid request, please verify your parameters."
	Message string `json:"message" example:"Invalid request, please verify your parameters."`
}
