package errors

import (
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/monitoring"
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		monitoring.GetTracer().NewError(err).Send()
		log.Fatalf("%s: %s", msg, err)
	}
}
