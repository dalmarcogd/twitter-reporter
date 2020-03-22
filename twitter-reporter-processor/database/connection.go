package database

import (
	"context"
	"fmt"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/environments"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.elastic.co/apm/module/apmgorm"
	_ "go.elastic.co/apm/module/apmgorm/dialects/postgres"
	"log"
	"sync"
	"time"
)

var (
	doOnce sync.Once
	db     *gorm.DB
)

func GetConnection(ctx context.Context) *gorm.DB {
	doOnce.Do(func() {
		env := environments.GetEnvironment()
		conn, err := apmgorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", env.PostgresURL, env.PostgresPort, env.PostgresUsername, env.PostgresDatabase, env.PostgresPassword))
		if err != nil {
			log.Fatal(err)
		}
		db = conn
		db.DB().SetConnMaxLifetime(time.Minute * 5)
		db.DB().SetMaxIdleConns(5)
		db.DB().SetMaxOpenConns(5)
	})
	return apmgorm.WithContext(ctx, db)
}

func CloseConnection() {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
