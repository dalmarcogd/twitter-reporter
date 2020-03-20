package main

import (
	"fmt"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/cache"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/database"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/environments"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/errors"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/handlers"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/middlewares"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.elastic.co/apm"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	database.Migrate()
	defer database.CloseConnection()
	defer cache.CloseConnection()
	env := environments.GetEnvironment()

	e := echo.New()
	e.Validator = utils.NewCustomValidator(validator.New())
	e.HTTPErrorHandler = errors.HttpErrorHandler()
	tracer, err := apm.NewTracer("twitter-reporter-api", "1.0.0")
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.Use(middlewares.ElasticApmMiddleware(middlewares.WithTracer(tracer)))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handlers.RegisterHandlers(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", env.Port)))
}
