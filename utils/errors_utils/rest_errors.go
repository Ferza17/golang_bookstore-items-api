package errors_utils

import (
	"net/http"
)

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad_Request",
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Not_Found",
	}
}

func NewUnauthorizedError() *RestError  {
	return &RestError{
		Message: "unable to retrieve user information given access_token",
		Status:  http.StatusUnauthorized,
		Error:   "Unauthorized",
	}
}

func NewInternalServerError(message string) *RestError  {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
