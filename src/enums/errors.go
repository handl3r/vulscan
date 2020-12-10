package enums

import (
	"fmt"
	"net/http"
)

type Error interface {
	error
	GetHttpCode() int
	GetCode() string
	GetMessage() string
}

type CustomError struct {
	HttpCode int
	Code     string
	Message  string
}

func NewCustomError(code, message string) *CustomError {
	return &CustomError{Code: code, Message: message}
}

func NewHttpCustomError(statusCode int, code, message string) *CustomError {
	return &CustomError{
		HttpCode: statusCode,
		Code:     code,
		Message:  message,
	}
}

func (c CustomError) Error() string {
	return fmt.Sprintf("[%s] %s", c.Code, c.Message)
}

func (c CustomError) GetHttpCode() int {
	return c.HttpCode
}

func (c CustomError) GetCode() string {
	return c.Code
}

func (c CustomError) GetMessage() string {
	return c.Message
}

var ErrEntityNotFound = NewCustomError("entity_not_found", "Entity not found")
var ErrResourceNotFound = NewHttpCustomError(
	http.StatusNotFound,
	"not_found",
	"Resource not found",
)
var ErrInvalidRequest = NewHttpCustomError(
	http.StatusBadRequest,
	"invalid_request",
	"Invalid request",
)
var ErrUnauthorized = NewHttpCustomError(
	http.StatusUnauthorized,
	"unauthorized",
	"Unauthorized",
)
var Forbidden = NewHttpCustomError(
	http.StatusForbidden,
	"forbidden",
	"Forbidden",
)
var ErrSystem = NewHttpCustomError(
	http.StatusInternalServerError,
	"system error",
	"System error. Please contact administrator",
)
var ErrNoResources = NewHttpCustomError(
	http.StatusOK,
	"no_resources",
	"Create your first resource",
)
