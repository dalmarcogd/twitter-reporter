package database

import (
	"context"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/errors"
	"net/http"
)

func GetReporterById(ctx context.Context, reporterId string) (ReporterModel, error) {
	reporter := ReporterModel{}
	GetConnection(ctx).Where("id = ?", reporterId).First(&reporter)
	if reporter.Id != "" {
		return reporter, nil
	}
	return reporter, errors.NewError(http.StatusNotFound, "Reporter not found", nil)
}
