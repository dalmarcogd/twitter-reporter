package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func HealthCheckHandler(c echo.Context) error {
	return c.String(http.StatusOK, "")
}
