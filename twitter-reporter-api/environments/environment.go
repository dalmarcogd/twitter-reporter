package environments

import (
	"github.com/crgimenes/goconfig"
	"log"
	"sync"
)

var (
	doOnce sync.Once
	env    Environment
)

type Environment struct {
	Port             string `cfg:"PORT" cfgDefault:"8000"`
	Environment      string `cfg:"ENVIRONMENT" cfgDefault:"transactions-dev"`
	RabbitURL        string `cfg:"RABBIT_URL" cfgDefault:"localhost"`
	RabbitPort       string `cfg:"RABBIT_AMQP_PORT" cfgDefault:"5672"`
	RabbitUsername   string `cfg:"RABBIT_USERNAME" cfgDefault:"rabbitmq"`
	RabbitPassword   string `cfg:"RABBIT_PASSWORD" cfgDefault:"rabbitmq"`
	RabbitVHost      string `cfg:"RABBIT_VHOST" cfgDefault:"/"`
	RedisHost        string `cfg:"REDIS_HOST" cfgDefault:"localhost"`
	RedisPort        string `cfg:"REDIS_PORT" cfgDefault:"6379"`
	PostgresURL      string `cfg:"POSTGRES_URL" cfgDefault:"localhost"`
	PostgresPort     string `cfg:"POSTGRES_PORT" cfgDefault:"5432"`
	PostgresUsername string `cfg:"POSTGRES_USERNAME" cfgDefault:"postgres"`
	PostgresPassword string `cfg:"POSTGRES_PASSWORD" cfgDefault:"postgres"`
	PostgresDatabase string `cfg:"POSTGRES_DATABASE" cfgDefault:"twitter"`
}

func GetEnvironment() Environment {
	doOnce.Do(func() {
		env = Environment{}
		err := goconfig.Parse(&env)
		if err != nil {
			log.Fatal(err)
		}
	})
	return env
}
