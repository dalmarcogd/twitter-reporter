package errors

import (
	"fmt"
)

type Error struct {
	StatusCode  int
	Description string
	Internal    error
}

func NewError(statusCode int, description string, internal error) *Error {
	return &Error{StatusCode: statusCode, Description: description, Internal: internal}
}

func (e Error) Error() string {
	return fmt.Sprintf("%v - %v = %v", e.StatusCode, e.Description, e.Internal)
}
