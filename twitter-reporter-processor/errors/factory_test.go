package errors

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestError_Error(t *testing.T) {
	assert.NotEmpty(t, NewError(http.StatusInternalServerError, "error", nil).Error())
}

func TestNewError(t *testing.T) {
	err := NewError(http.StatusInternalServerError, "error", nil)
	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, "error", err.Description)
}
