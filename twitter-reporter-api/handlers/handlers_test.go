package handlers

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterHandlers(t *testing.T) {
	e := echo.New()
	RegisterHandlers(e)
	assert.GreaterOrEqual(t, len(e.Routes()), 2)
}
