package main

import (
	"fmt"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/cache"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/database"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/environments"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/errors"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/handlers"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/middlewares"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/monitoring"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	e := echo.New()
	e.Validator = utils.NewCustomValidator(validator.New())
	e.HTTPErrorHandler = errors.HttpErrorHandler()
	e.Renderer = utils.NewTemplateRenderer()

	errors.FailOnError(database.Migrate(), "Fail on migrate database")

	defer database.CloseConnection()
	defer cache.CloseConnection()

	e.Use(middlewares.ElasticApmMiddleware(middlewares.WithTracer(monitoring.GetTracer())))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handlers.RegisterHandlers(e)

	errors.FailOnError(e.Start(fmt.Sprintf(":%s", environments.GetEnvironment().Port)), "Fail on start http server")
}
