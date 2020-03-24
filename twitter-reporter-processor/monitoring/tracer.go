package monitoring

import (
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/environments"
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
		t, err := apm.NewTracer(env.ServiceName, env.ServiceVersion)
		if err != nil {
			log.Fatal(err)
		}
		tracer = t
	})
	return tracer
}
