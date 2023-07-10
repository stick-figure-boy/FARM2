package apierror

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	badRequest          = "bad request"
	unauthorized        = "unauthorized"
	forbidden           = "forbidden"
	notFound            = "not found"
	conflict            = "conflict"
	internalServerError = "internal server error"
)

var (
	ErrBadRequest          = errors.New(badRequest)
	ErrUnauthorized        = errors.New(unauthorized)
	ErrForbidden           = errors.New(forbidden)
	ErrNotFound            = errors.New(notFound)
	ErrConflict            = errors.New(conflict)
	ErrInternalServerError = errors.New(internalServerError)
)

type ErrorResponse struct {
	ErrorCode string `json:"error_code"`
	Error     string `json:"error"`
	Message   string `json:"message"`
}

type APIError struct {
	StatusCode int
	ErrorCode  string
	Error      string
	Message    string
}

func NewAPIError(err error, code string, msg string) *APIError {
	e := APIError{
		ErrorCode: code,
		Error:     err.Error(),
		Message:   msg,
	}

	switch err {
	case ErrBadRequest:
		e.StatusCode = http.StatusBadRequest
	case ErrUnauthorized:
		e.StatusCode = http.StatusUnauthorized
	case ErrForbidden:
		e.StatusCode = http.StatusForbidden
	case ErrNotFound:
		e.StatusCode = http.StatusNotFound
	case ErrConflict:
		e.StatusCode = http.StatusConflict
	case ErrInternalServerError:
		e.StatusCode = http.StatusInternalServerError
	default:
		e.StatusCode = http.StatusInternalServerError
		e.ErrorCode = UnknownErrCode
		e.Error = internalServerError
		e.Message = "Unknown error occurred."
	}

	return &e
}

func (e *APIError) ResponseParser(ctx *gin.Context) {
	ctx.JSON(e.StatusCode, ErrorResponse{
		ErrorCode: e.ErrorCode,
		Error:     e.Error,
		Message:   e.Message,
	})
}
