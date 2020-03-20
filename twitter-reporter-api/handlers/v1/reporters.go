package v1

import (
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/brokers/events"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/brokers/rabbit"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
)

type ReportersRequest struct {
	Tag string `json:"tag" validate:"required"`
}

type ReportersResponse struct {
	ReporterId string `json:"reporter_id" validate:"required"`
	Tag        string `json:"tag" validate:"required"`
}

func ReportersV1Handler(c echo.Context) error {
	reporterRequest := new(ReportersRequest)
	if err := c.Bind(reporterRequest); err != nil {
		return err
	}

	if err := c.Validate(reporterRequest); err != nil {
		return err
	}

	uid, _ := uuid.NewUUID()
	event := events.NewReporterEvent(uid.String(), reporterRequest.Tag)

	if err := rabbit.NewRabbit().Publish(event); err != nil {
		return err
	}

	reporterResponse := new(ReportersResponse)
	reporterResponse.ReporterId = event.ReporterId
	reporterResponse.Tag = event.Tag
	return c.JSON(http.StatusOK, reporterResponse)
}
