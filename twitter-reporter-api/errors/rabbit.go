package errors

import "net/http"

func NewRabbitConnectionError(err ...error) *Error {
	var erro error
	if len(err) > 0 {
		erro = err[0]
	}

	return NewError(http.StatusBadRequest, "Error when connect to Rabbit.", erro)
}
