package wrapper

import (
	"net/http"
)

// Response is a generic API response structure
type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

// NewResponse creates a new response with default values
func NewResponse() *Response {
	return &Response{
		Success:    false,
		StatusCode: http.StatusOK,
		Message:    "",
	}
}

// SuccessResponse returns a success response with data
func SuccessResponse(data interface{}, message string) *Response {
	return &Response{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    message,
		Data:       data,
	}
}

// ErrorResponse returns an error response
func ErrorResponse(err interface{}, message string, statusCode int) *Response {
	return &Response{
		Success:    false,
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
	}
}

// NotFoundResponse returns a not found error response
func NotFoundResponse(message string) *Response {
	if message == "" {
		message = "Resource not found"
	}
	return ErrorResponse(nil, message, http.StatusNotFound)
}

// BadRequestResponse returns a bad request error response
func BadRequestResponse(err interface{}, message string) *Response {
	if message == "" {
		message = "Bad request"
	}
	return ErrorResponse(err, message, http.StatusBadRequest)
}

// ServerErrorResponse returns a server error response
func ServerErrorResponse(err interface{}, message string) *Response {
	if message == "" {
		message = "Internal server error"
	}
	return ErrorResponse(err, message, http.StatusInternalServerError)
}

// UnauthorizedResponse returns an unauthorized error response
func UnauthorizedResponse(message string) *Response {
	if message == "" {
		message = "Unauthorized access"
	}
	return ErrorResponse(nil, message, http.StatusUnauthorized)
}
