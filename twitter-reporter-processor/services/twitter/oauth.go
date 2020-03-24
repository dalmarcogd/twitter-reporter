package twitter

import (
	"context"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/errors"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/utils"
	"go.elastic.co/apm/module/apmhttp"
	"io/ioutil"
	"net/http"
)

func GetBearerToken(ctx context.Context) (string, error) {
	req, err := http.NewRequest(http.MethodPost, "https://api.twitter.com/oauth2/token?grant_type=client_credentials", nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth("xEZ4cABBOC1BXd0DtAvFSzNin", "7A4Rt4kHN5nP1ydPIUY0M5xH23b2TWEjKUC6DlMBvMUTRFImIT")

	client := apmhttp.WrapClient(http.DefaultClient)
	response, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", errors.NewError("Error when get bearer token from twitter", err)
	}

	body := make(map[string]string)
	bodyData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	if err := utils.NewJsonConverter().Decode(bodyData, &body); err != nil {
		return "", err
	}

	if bearerToken, ok := body["access_token"]; ok {
		return bearerToken, nil
	}
	return "", errors.NewError("Error when get bearer token from twitter", err)
}
