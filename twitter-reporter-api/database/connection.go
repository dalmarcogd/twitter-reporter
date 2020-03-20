package database

import (
	"fmt"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/environments"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"sync"
	"time"
)

var (
	doOnce sync.Once
	db *gorm.DB
)

func GetConnection() *gorm.DB {
	doOnce.Do(func() {
		env := environments.GetEnvironment()
		conn, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", env.PostgresURL, env.PostgresPort, env.PostgresUsername, env.PostgresDatabase, env.PostgresPassword))
		if err != nil {
			log.Fatal(err)
		}
		db = conn
		db.DB().SetConnMaxLifetime(time.Minute*5)
		db.DB().SetMaxIdleConns(5)
		db.DB().SetMaxOpenConns(5)
	})
	return db
}

func CloseConnection() {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
