package httpx

import (
	"net/http"
)

type ApiError struct {
	Status  int    `json:"-"`
	Code    string `json:"code"` // NOT_FOUND, VALIDATION_ERROR
	Message string `json:"message"`
	Details any    `json:"details,omitempty"` // validation errors
	Err     error  `json:"-"`                 // for dev logs
}

func (e *ApiError) Error() string { return e.Message }
func (e *ApiError) Unwrap() error { return e.Err }

func BadRequest(msg string, details any) *ApiError {
	return &ApiError{
		Status:  http.StatusBadRequest,
		Code:    "BAD_REQUEST",
		Message: msg,
		Details: details,
	}
}

func NotFound(msg string) *ApiError {
	return &ApiError{
		Status:  http.StatusNotFound,
		Code:    "NOT_FOUND",
		Message: msg,
	}
}

func Internal(err error) *ApiError {
	return &ApiError{
		Status:  http.StatusInternalServerError,
		Code:    "INTERNAL",
		Message: "Something went wrong",
		Err:     err,
	}
}
