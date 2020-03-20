package environments

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEnvironment(t *testing.T) {
	env := GetEnvironment()
	assert.NotNil(t, env.Environment)
}
