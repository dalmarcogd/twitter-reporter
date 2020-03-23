package database

import (
	"context"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/errors"
	"net/http"
	"time"
)

func GetReporterById(ctx context.Context, reporterId string) (ReporterModel, error) {
	reporter := ReporterModel{}
	GetConnection(ctx).Where("id = ?", reporterId).First(&reporter)
	if reporter.Id != "" {
		return reporter, nil
	}
	return reporter, errors.NewError(http.StatusNotFound, "Reporter not found", nil)
}

func GetTwitterTweets(ctx context.Context, day time.Time) ([]TwitterTweetModel, error) {
	tweets := make([]TwitterTweetModel, 0)
	day1 := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
	day2 :=  day1.Add(time.Hour * 24)
	db := GetConnection(ctx).Model(&TwitterTweetModel{}).Where("created_at BETWEEN ? AND ?", day1, day2).Find(&tweets)
	return tweets, db.Error
}

func GetReporterTopFollowersUserById(ctx context.Context, reporterId string) ([]map[string]interface{}, error) {
	results := make([]map[string]interface{}, 0)

	rows, err := GetConnection(ctx).
		Table("tweets").
		Select("users.name, count(*)").
		Joins("join users on users.id = tweets.twitter_user").
		Where("tweets.reporter = ?", reporterId).
		Group("users.name").
		Order("count(*) desc").
		Limit(5).Rows()

	if err != nil {
		return results, err
	}
	defer rows.Close()
	var name string
	var count int
	for rows.Next() {
		err := rows.Scan(&name, &count)
		if err != nil {
			return results, err
		}
		results = append(results, map[string]interface{}{"name": name, "count": count})
	}

	return results, db.Error
}

func GetReporterTweetsHour(ctx context.Context) (map[int]int, error) {
	results := make(map[int]int, 0)

	tweets, err := GetTwitterTweets(ctx, time.Now())
	if err != nil {
		return results, err
	}

	for _, tweet := range tweets {
		hour := tweet.CreatedAt.Hour()
		results[hour] = results[hour] + 1
	}

	return results, db.Error
}


func GetReporterTweetsLanguagesCountries(ctx context.Context) (map[string]map[string]int, error) {
	results := make(map[string]map[string]int, 0)

	tweets, err := GetTwitterTweets(ctx, time.Now())
	if err != nil {
		return results, err
	}

	reporterMap := make(map[string]ReporterModel)

	for _, tweet := range tweets {
		reporterId := tweet.Reporter

		var tag string
		if reporter, ok := reporterMap[reporterId]; ok{
			tag = reporter.Tag
		} else {
			reporter, err := GetReporterById(ctx, reporterId)
			if err != nil {
				return results, err
			}
			tag = reporter.Tag
			reporterMap[tag] = reporter
		}
		if _, ok := results[tag]; !ok {
			results[tag] = map[string]int{}
		}
		results[tag][tweet.Language] += 1
	}

	return results, db.Error
}