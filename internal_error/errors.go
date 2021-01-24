package internal_error

import (
	"encoding/json"
)

const (
	JsonParseError      = "JsonParseError"
	ValidationError     = "ValidationError"
	InternalServerError = "InternalServerError"
)

type ErrorResponse struct {
	StatusCode       int    `json:"statusCode"`
	ErrorName        string `json:"errorName"`
	ErrorDescription string `json:"errorDescription"`
}

func (e *ErrorResponse) Error() string {
	errorResponse, _ := json.Marshal(e)
	return string(errorResponse)
}

func CreateJsonParseError(err error) error {
	return &ErrorResponse{
		StatusCode:       400,
		ErrorName:        JsonParseError,
		ErrorDescription: err.Error(),
	}
}

func CreateValidationError(err error) error {
	return &ErrorResponse{
		StatusCode:       400,
		ErrorName:        ValidationError,
		ErrorDescription: err.Error(),
	}
}
func CreateInternalServerError(err error) error {
	return &ErrorResponse{
		StatusCode:       500,
		ErrorName:        InternalServerError,
		ErrorDescription: err.Error(),
	}
}