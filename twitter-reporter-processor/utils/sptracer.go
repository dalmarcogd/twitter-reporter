package utils

import (
	"context"
	"go.elastic.co/apm"
)

func SpanTracer(ctx context.Context, name string, spanType string, f func(cx context.Context, span *apm.Span) error) error {
	tx := apm.TransactionFromContext(ctx)
	span := tx.StartSpan(name, spanType, apm.SpanFromContext(ctx))
	defer span.End()
	err := f(apm.ContextWithSpan(ctx, span), span)
	return err
}
