package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRabbitConnectionError(t *testing.T) {
	err := NewRabbitConnectionError()
	assert.Equal(t, "Error when connect to Rabbit.", err.Description)
}
