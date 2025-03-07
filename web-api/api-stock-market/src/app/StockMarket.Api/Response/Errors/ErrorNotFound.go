package Errors

// ErrorNotFound represents the 404 error.
// swagger:model
type ErrorNotFound struct {
	// Error code.
	// example: 404
	Code int `json:"code" example:"404"`
	// Error description.
	// example: "Not found"
	Error string `json:"error" example:"Not found"`
	// Error message.
	// example: "Resource not found."
	Message string `json:"message" example:"Resource not found."`
}
