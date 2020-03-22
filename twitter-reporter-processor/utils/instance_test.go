package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsInstanceOf(t *testing.T) {
	assert.True(t, IsInstanceOf(nil, nil))
	assert.False(t, IsInstanceOf("", nil))
	assert.False(t, IsInstanceOf(nil, ""))
	assert.True(t, IsInstanceOf("", ""))
}