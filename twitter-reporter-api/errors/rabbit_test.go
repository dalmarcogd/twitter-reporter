package errors

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewRabbitConnectionError(t *testing.T) {
	err := NewRabbitConnectionError()
	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
}
