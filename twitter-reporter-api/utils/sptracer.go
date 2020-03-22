package utils

import (
	"context"
	"go.elastic.co/apm"
)

func SpanTracer(ctx context.Context, name string, spanType string, f func(cx context.Context, span *apm.Span) error) error {
	tx := apm.TransactionFromContext(ctx)
	span := tx.StartSpan(name, spanType, nil)
	defer span.End()
	err := f(ctx, span)
	return err
}
