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
	pedidosGroupV1.POST("/reporters", v1.ReportersPostV1Handler)
	pedidosGroupV1.GET("/reporters/:reporterId/top/followers/users", v1.ReportersTopFollowersUserGetV1Handler)
	pedidosGroupV1.GET("/reporters/tweets/hours", v1.ReportersTweetsHoursGetV1Handler)
	pedidosGroupV1.GET("/reporters/tweets/languages/countries", v1.ReportersTweetsLanguagesCountriesGetV1Handler)
}
