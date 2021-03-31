package errs

import "net/http"

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:",omitempty"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{Message: e.Message}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{message, http.StatusNotFound}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{message, http.StatusInternalServerError}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{message, http.StatusBadRequest}
}

func NewValidationError(message string) *AppError {
	return &AppError{message, http.StatusUnprocessableEntity}
}
