package database

import (
	"context"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/errors"
	"time"
)

func GetReporterById(ctx context.Context, reporterId string) (ReporterModel, error) {
	reporter := ReporterModel{}
	GetConnection(ctx).Where("id = ?", reporterId).First(&reporter)
	if reporter.Id != "" {
		return reporter, nil
	}
	return reporter, errors.NewError("Reporter not found", nil)
}

func CreateReporter(ctx context.Context, reporterId string, tag string) error {
	return GetConnection(ctx).Save(&ReporterModel{Id: reporterId, Tag: tag}).Error
}

func GetTwitterUserById(ctx context.Context, twitterUserId float64) (TwitterUserModel, error) {
	twitterUser := TwitterUserModel{}
	GetConnection(ctx).Where("id = ?", twitterUserId).First(&twitterUser)
	if twitterUser.Id != 0 {
		return twitterUser, nil
	}
	return twitterUser, errors.NewError("User not found", nil)
}

func CreateTwitterUser(ctx context.Context, twitterUserId float64, name string, screenName string, statusesCount float64, followersCount float64, location string) (TwitterUserModel, error) {
	twitterUser := TwitterUserModel{Id: twitterUserId, Name: name, ScreenName: screenName, StatusesCount: statusesCount, FollowersCount: followersCount, Location: location}
	return twitterUser, GetConnection(ctx).Save(&twitterUser).Error
}

func UpdateTwitterUser(ctx context.Context, twitterUserId float64, name string, screenName string, statusesCount float64, followersCount float64, location string) (TwitterUserModel, error) {
	twitterUser := TwitterUserModel{Id: twitterUserId, Name: name, ScreenName: screenName, StatusesCount: statusesCount, FollowersCount: followersCount, Location: location}
	return twitterUser, GetConnection(ctx).Model(&twitterUser).Update(&twitterUser).Error
}

func GetTwitterTweetById(ctx context.Context, twitterTweetId float64) (TwitterTweetModel, error) {
	twitterTweet := TwitterTweetModel{}
	GetConnection(ctx).Where("id = ?", twitterTweetId).First(&twitterTweet)
	if twitterTweet.Id != 0 {
		return twitterTweet, nil
	}
	return twitterTweet, errors.NewError("Tweet not found", nil)
}

func CreateTwitterTweet(ctx context.Context, twitterTweetId float64, reporter string, twitterUser float64, text string, language string, createdAt time.Time) error {
	return GetConnection(ctx).Save(&TwitterTweetModel{Id: twitterTweetId, Reporter: reporter, TwitterUser: twitterUser, Text: text, Language: language, CreatedAt: createdAt}).Error
}
