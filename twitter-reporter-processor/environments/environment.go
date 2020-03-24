package environments

import (
	"github.com/crgimenes/goconfig"
	"log"
	"os"
	"sync"
)

var (
	doOnce sync.Once
	env    Environment
)

type Environment struct {
	Environment      string `cfg:"ENVIRONMENT" cfgDefault:"local"`
	ServiceName      string `cfg:"SERVICE_NAME" cfgDefault:"twitter-reporter-processor"`
	ServiceVersion   string `cfg:"SERVICE_VERSION" cfgDefault:"1.0.0"`
	RabbitURL        string `cfg:"RABBIT_URL" cfgDefault:"localhost"`
	RabbitPort       string `cfg:"RABBIT_AMQP_PORT" cfgDefault:"5672"`
	RabbitUsername   string `cfg:"RABBIT_USERNAME" cfgDefault:"rabbitmq"`
	RabbitPassword   string `cfg:"RABBIT_PASSWORD" cfgDefault:"rabbitmq"`
	RabbitVHost      string `cfg:"RABBIT_VHOST" cfgDefault:"/"`
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
		customEnvironment(env)
		if err != nil {
			log.Fatal(err)
		}
	})
	return env
}

func customEnvironment(env Environment) {
	_ = os.Setenv("ELASTIC_APM_SERVICE_NAME", env.ServiceName)
	_ = os.Setenv("ELASTIC_APM_SERVICE_VERSION", env.ServiceVersion)
	_ = os.Setenv("ELASTIC_APM_ENVIRONMENT", env.Environment)
}
