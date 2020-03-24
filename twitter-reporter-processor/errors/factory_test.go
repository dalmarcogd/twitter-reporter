package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError_Error(t *testing.T) {
	assert.NotEmpty(t, NewError( "error", nil).Error())
}

func TestNewError(t *testing.T) {
	err := NewError( "error", nil)
	assert.Equal(t, "error", err.Description)
}
