package errors

import (
	"context"
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-processor/monitoring"
	"go.elastic.co/apm"
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		monitoring.GetTracer().NewError(err).Send()
		log.Fatalf("%s: %s", msg, err)
	}
}

func PrintOnError(ctx context.Context, err error, msg string) {
	if err != nil {
		apmError := monitoring.GetTracer().NewError(err)
		apmError.SetTransaction(apm.TransactionFromContext(ctx))
		apmError.SetSpan(apm.SpanFromContext(ctx))
		apmError.Send()
		log.Printf("%s: %s", msg, err)
	}
}
