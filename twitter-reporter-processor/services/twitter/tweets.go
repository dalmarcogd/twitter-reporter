package twitter

import (
	"context"
	"fmt"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/utils"
	"go.elastic.co/apm/module/apmhttp"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetTweetsByHashtag(ctx context.Context, hashtag string) ([]interface{}, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.twitter.com/1.1/search/tweets.json?result_type=recent&count=100&q="+strings.ReplaceAll(hashtag, "#", "%23"), nil)
	if err != nil {
		return make([]interface{}, 0), err
	}
	bearerToken, err := GetBearerToken(ctx)
	if err != nil {
		return make([]interface{}, 0), err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
	client := apmhttp.WrapClient(http.DefaultClient)
	response, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return make([]interface{}, 0), err
	}
	body := make(map[string]interface{})
	bodyData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return make([]interface{}, 0), err
	}
	if err := utils.NewJsonConverter().Decode(bodyData, &body); err != nil {
		return make([]interface{}, 0), err
	}
	return body["statuses"].([]interface{}), nil
}
