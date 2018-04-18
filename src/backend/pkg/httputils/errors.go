package httputils

import (
	"net/http"
	"encoding/json"
)

type ErrNotFound interface {
	error
	NotFound()
}

type errNotFound struct {
	Err error
}

func (errNotFound) NotFound() {

}

// HTTPError represents an HTTP error with HTTP status code and error message
type HTTPError interface {
	error
	// StatusCode returns the HTTP status code of the error
	StatusCode() int
	Headers() http.Header
	WithError(err error) HTTPError
	InsideError() error
}

// apiError represents an error that can be sent in an error response.
type APIError struct {
	// Status represents the HTTP status code
	status int `json:"-"`
	// ErrorCode is the code uniquely identifying an error
	// ErrorCode string `json:"error_code"`
	// Message is the error message that may be displayed to end users
	Message      string `json:"message"`
	DebugMessage string `json:"debug_message,omitempty"`
	// Details specifies the additional error information
	Errors interface{} `json:"errors,omitempty"`
	err    error
}

func (e *APIError) WithError(err error) HTTPError {
	e.err = err
	return e
}

func (e *APIError) InsideError() error {
	return e.err
}

// Error returns the error message.
func (e *APIError) Error() string {
	//if jsonData, err := e.ToJson(); err == nil {
	//	return string(jsonData)
	//} else {
	//	return err.Error()
	//}
	// TODO: error handing
	b, _ := json.Marshal(e)
	return string(b)
}

// StatusCode returns the HTTP status code.
func (e *APIError) StatusCode() int {
	return e.status
}

func (e *APIError) Headers() http.Header {
	h := http.Header{}
	h.Add("Content-Type", "application/json")
	return h
}

func NewAPIError(status int, message, debugMessage string, errors ...interface{}) *APIError {
	apiError := &APIError{
		status:  status,
		Message: message,
		// ErrorCode: errorCode,
		DebugMessage: debugMessage,
	}
	if len(errors) > 0 {
		apiError.Errors = errors[0]
	}
	return apiError
}

// InternalServerError creates a new API error representing an internal server error (HTTP 500)
func InternalServerError(message string, debugMessage ...string) *APIError {

	if message == "" {
		message = http.StatusText(http.StatusInternalServerError)
	}

	debugMsg := ""

	if len(debugMessage) > 0 {
		debugMsg = debugMessage[0]
	}
	return NewAPIError(http.StatusInternalServerError, message, debugMsg)
}

// NotFound creates a new API error representing a resource-not-found error (HTTP 404)
func NotFound(message string, debugMessage ...string) *APIError {
	if message == "" {
		message = http.StatusText(http.StatusNotFound)
	}

	debugMsg := ""

	if len(debugMessage) > 0 {
		debugMsg = debugMessage[0]
	}
	return NewAPIError(http.StatusNotFound, message, debugMsg)
}

// Unauthorized creates a new API error representing an authentication failure (HTTP 401)
func Unauthorized(message ...string) *APIError {
	var msg string
	if len(message) > 0 {
		msg = message[0]
	} else {
		msg = http.StatusText(http.StatusUnauthorized)
	}
	return NewAPIError(http.StatusUnauthorized, msg, "")
}

func Forbidden(message ...string) *APIError {
	var msg string
	if len(message) > 0 {
		msg = message[0]
	} else {
		msg = http.StatusText(http.StatusForbidden)
	}
	return NewAPIError(http.StatusForbidden, msg, "")
}

func BadRequest(message ...string) *APIError {
	var msg string
	if len(message) > 0 {
		msg = message[0]
	} else {
		msg = http.StatusText(http.StatusBadRequest)
	}
	return NewAPIError(http.StatusBadRequest, msg, "")
}
