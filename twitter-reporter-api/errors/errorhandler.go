package errors

import (
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/monitoring"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/utils"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func HttpErrorHandler() echo.HTTPErrorHandler {
	return func(err error, context echo.Context) {
		var status int
		monitoring.GetTracer().NewError(err).Send()
		if utils.IsInstanceOf(err, &echo.HTTPError{}) {
			status = err.(*echo.HTTPError).Code
		} else if utils.IsInstanceOf(err, &Error{}) {
			status = err.(*Error).StatusCode
		} else if utils.IsInstanceOf(err, validator.ValidationErrors{}) {
			status = http.StatusUnprocessableEntity
		}

		_ = context.JSON(status, err)
	}
}
