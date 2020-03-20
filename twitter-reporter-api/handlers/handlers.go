package handlers

import (
	v1 "github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/handlers/v1"
	"github.com/labstack/echo"
)

func RegisterHandlers(e *echo.Echo) {
	e.GET("/health-check", HealthCheckHandler)
	// Reporters
	pedidosGroup := e.Group("/reporter-api")
	pedidosGroupV1 := pedidosGroup.Group("/v1")
	pedidosGroupV1.POST("/reporters", v1.ReportersV1Handler)
}
