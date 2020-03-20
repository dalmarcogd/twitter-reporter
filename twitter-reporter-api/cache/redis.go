package cache

import (
	"fmt"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/environments"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/xuyu/goredis"
	"log"
	"sync"
)

var (
	doOnce sync.Once
	client *goredis.Redis
)

func GetConnection() *goredis.Redis {
	doOnce.Do(func() {
		env := environments.GetEnvironment()
		conn, err := goredis.Dial(&goredis.DialConfig{Address: fmt.Sprintf("%s:%s", env.RedisHost, env.RedisPort)})

		if err != nil {
			log.Fatal(err)
		}
		client = conn
	})
	return client
}

func CloseConnection() {
	client.ClosePool()
}
