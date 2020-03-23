package v1

import (
	"context"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/brokers/events"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/brokers/rabbit"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/database"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/errors"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"go.elastic.co/apm"
	"net/http"
	"sort"
)

type ReportersRequest struct {
	Tag string `json:"tag" validate:"required"`
}

type ReportersResponse struct {
	ReporterId string `json:"reporter_id" validate:"required"`
	Tag        string `json:"tag" validate:"required"`
}

func ReportersPostV1Handler(c echo.Context) error {
	reporterRequest := new(ReportersRequest)
	if err := c.Bind(reporterRequest); err != nil {
		return err
	}

	if err := c.Validate(reporterRequest); err != nil {
		return err
	}

	uid, _ := uuid.NewUUID()
	ctx := c.Request().Context()
	var event *events.ReporterEvent

	if err := utils.SpanTracer(ctx, "Publish on rabbit", "amqp", func(ctx context.Context, span *apm.Span) error {
		event = events.NewReporterEvent(ctx, uid.String(), reporterRequest.Tag)
		return rabbit.Publish(ctx, event)
	}); err != nil {
		return err
	}

	reporterResponse := new(ReportersResponse)
	reporterResponse.ReporterId = event.ReporterId
	reporterResponse.Tag = event.Tag
	return c.JSON(http.StatusOK, reporterResponse)
}

type ReportersTopFollowersUserResult struct {
	Name  string `json:"name" validate:"required"`
	Count int    `json:"count" validate:"required"`
}

type ReportersTopFollowersUserResponse struct {
	ReporterId string                            `json:"reporter_id" validate:"required"`
	Tag        string                            `json:"tag" validate:"required"`
	Results    []ReportersTopFollowersUserResult `json:"results" validate:"required"`
}

func ReportersTopFollowersUserGetV1Handler(c echo.Context) error {
	reporterId := c.Param("reporterId")
	if reporterId == "" {
		return errors.NewError(http.StatusUnprocessableEntity, "Param reporterId is required.", nil)
	}
	var reporter database.ReporterModel
	var err error
	if err := utils.SpanTracer(c.Request().Context(), "Get Reporter", "database.queries", func(cx context.Context, span *apm.Span) error {
		reporter, err = database.GetReporterById(c.Request().Context(), reporterId)
		return err
	}); err != nil {
		return err
	}

	var results []map[string]interface{}
	if err := utils.SpanTracer(c.Request().Context(), "Get Top Followers by User", "database.queries", func(cx context.Context, span *apm.Span) error {
		results, err = database.GetReporterTopFollowersUserById(c.Request().Context(), reporterId)
		return err
	}); err != nil {
		return err
	}

	reporterResponse := new(ReportersTopFollowersUserResponse)
	for _, result := range results {
		reporterResponse.Results = append(reporterResponse.Results, ReportersTopFollowersUserResult{Name: result["name"].(string), Count: result["count"].(int)})
	}
	reporterResponse.ReporterId = reporter.Id
	reporterResponse.Tag = reporter.Tag

	return c.JSON(http.StatusOK, reporterResponse)
}

type ReportersTweetsHoursResult struct {
	Hour  int `json:"hour" validate:"required"`
	Count int `json:"count" validate:"required"`
}

type ReportersTweetsHoursResponse struct {
	Results []ReportersTweetsHoursResult `json:"results" validate:"required"`
}

func ReportersTweetsHoursGetV1Handler(c echo.Context) error {
	var results map[int]int
	var err error
	if err := utils.SpanTracer(c.Request().Context(), "Get Tweets Hours", "database.queries", func(cx context.Context, span *apm.Span) error {
		results, err = database.GetReporterTweetsHour(c.Request().Context())
		return err
	}); err != nil {
		return err
	}

	keys := make([]int, 0)
	for k := range results {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	reporterResponse := new(ReportersTweetsHoursResponse)
	for _, k := range keys {
		reporterResponse.Results = append(reporterResponse.Results, ReportersTweetsHoursResult{Hour: k, Count: results[k]})
	}

	return c.JSON(http.StatusOK, reporterResponse)
}

type ReportersTweetsLanguagesCountriesResult struct {
	Tag                  string         `json:"tag" validate:"required"`
	LanguageCountryCount map[string]int `json:"language_country_count" validate:"required"`
}

type ReportersTweetsLanguagesCountriesResponse struct {
	Results []ReportersTweetsLanguagesCountriesResult `json:"results" validate:"required"`
}

func ReportersTweetsLanguagesCountriesGetV1Handler(c echo.Context) error {
	var results map[string]map[string]int
	var err error
	if err := utils.SpanTracer(c.Request().Context(), "Get Tweets Languages/Countries", "database.queries", func(cx context.Context, span *apm.Span) error {
		results, err = database.GetReporterTweetsLanguagesCountries(c.Request().Context())
		return err
	}); err != nil {
		return err
	}

	keys := make([]string, 0)
	for k := range results {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	reporterResponse := new(ReportersTweetsLanguagesCountriesResponse)
	for _, k := range keys {
		reporterResponse.Results = append(reporterResponse.Results, ReportersTweetsLanguagesCountriesResult{Tag: k, LanguageCountryCount: results[k]})
	}

	return c.JSON(http.StatusOK, reporterResponse)
}
