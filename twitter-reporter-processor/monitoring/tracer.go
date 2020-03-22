package monitoring

import (
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
		t, err := apm.NewTracer("twitter-reporter-processor", "1.0.0")
		if err != nil {
			log.Fatal(err)
		}
		tracer = t
	})
	return tracer
}
