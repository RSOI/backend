package ui

import (
	"errors"

	"github.com/jackc/pgx"
)

// Response interface
type Response struct {
	Status int         `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

// CodeToError code -> message
func CodeToError(code int) string {
	switch code {
	case 400:
		return "Validation error"
	case 404:
		return "Requested data not found"
	case 500:
		return "Server error. Contact admins for additional info"
	case 0:
		return "One of microservice is unavailable"
	}

	return "OK"
}

var (
	// ErrNoResult - no data found
	ErrNoResult = errors.New("no data found")
	// ErrNoDataToDelete - no data found to delete"
	ErrNoDataToDelete = errors.New("no data found to delete")
	// ErrNoDataToUpdate - no data found to delete"
	ErrNoDataToUpdate = errors.New("no data found to update")
	// ErrUnavailable - database is unavailable
	ErrUnavailable = errors.New("database is unavailable")
	// ErrFieldsRequired some of required fields are missing
	ErrFieldsRequired = errors.New("missed required field(s)")
)

// ErrToResponse status -> error
func ErrToResponse(err error) (int, string) {
	var statusCode int
	var statusText string

	if err != nil {
		statusText = err.Error()
	}

	switch err {
	case nil:
		statusCode = 200
	case ErrFieldsRequired:
		statusCode = 400
	case pgx.ErrNoRows:
		statusText = ErrNoResult.Error()
		statusCode = 404
	case ErrNoResult:
		statusCode = 404
	case ErrNoDataToDelete:
		statusCode = 404
	case ErrNoDataToUpdate:
		statusCode = 404
	case ErrUnavailable:
		statusCode = 503
	default:
		statusCode = 500
		//statusText = "Server error. Additional information may be contained in server logs."
	}

	return statusCode, statusText
}
