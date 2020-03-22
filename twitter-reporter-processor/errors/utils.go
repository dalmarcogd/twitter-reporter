package errors

import (
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/monitoring"
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		monitoring.GetTracer().NewError(err).Send()
		log.Fatalf("%s: %s", msg, err)
	}
}

func PrintOnError(err error, msg string) {
	if err != nil {
		monitoring.GetTracer().NewError(err).Send()
		log.Printf("%s: %s", msg, err)
	}
}