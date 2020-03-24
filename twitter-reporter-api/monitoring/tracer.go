package monitoring

import (
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/environments"
	"go.elastic.co/apm"
	"log"
	"sync"
)

var (
	doOnce sync.Once
	tracer *apm.Tracer
)

func GetTracer() *apm.Tracer {
	doOnce.Do(func() {
		env := environments.GetEnvironment()
		t, err := apm.NewTracerOptions(apm.TracerOptions{
			ServiceName:    env.ServiceName,
			ServiceVersion: env.ServiceVersion,
			ServiceEnvironment: env.Environment,
		})
		if err != nil {
			log.Fatal(err)
		}
		tracer = t
	})
	return tracer
}
