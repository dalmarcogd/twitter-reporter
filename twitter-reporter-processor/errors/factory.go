package errors

import (
	"fmt"
)

type Error struct {
	Description string
	Internal    error
}

func NewError(description string, internal error) *Error {
	return &Error{Description: description, Internal: internal}
}

func (e Error) Error() string {
	return fmt.Sprintf("%v = %v", e.Description, e.Internal)
}
