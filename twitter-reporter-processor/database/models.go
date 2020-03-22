package database

import "time"

type ReporterModel struct {
	Id     string `gorm:"primary_key"`
	Tag    string
	Tweets []TwitterTweetModel `gorm:"foreignkey:Reporter;association_foreignkey:Id"`
}

func (ReporterModel) TableName() string {
	return "reporters"
}

type TwitterUserModel struct {
	Id             float64 `gorm:"primary_key"`
	Name           string
	ScreenName     string
	StatusesCount  float64
	FollowersCount float64
	Location       string
	Tweets         []TwitterTweetModel `gorm:"foreignkey:TwitterUser;association_foreignkey:Id"`
}

func (TwitterUserModel) TableName() string {
	return "users"
}

type TwitterTweetModel struct {
	Id          float64 `gorm:"primary_key"`
	Reporter    string
	TwitterUser float64
	Text        string
	Language    string
	CreatedAt   time.Time
}

func (TwitterTweetModel) TableName() string {
	return "tweets"
}
