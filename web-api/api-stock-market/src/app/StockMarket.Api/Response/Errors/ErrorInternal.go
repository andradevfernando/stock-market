package Errors

// ErrorInternal represents the 500 error.
// swagger:model
type ErrorInternal struct {
	// Error code.
	// example: 500
	Code int `json:"code" example:"500"`
	// Error description.
	// example: "Internal server error"
	Error string `json:"error" example:"Internal server error"`
	// Error message.
	// example: "Internal server error, please try again later."
	Message string `json:"message" example:"Internal server error, please try again later."`
}
