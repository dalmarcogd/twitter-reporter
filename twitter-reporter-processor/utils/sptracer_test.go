package utils

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.elastic.co/apm"
	"testing"
)

func TestSpanTracer(t *testing.T) {
	err := SpanTracer(context.Background(), "", "test", func(cx context.Context, span *apm.Span) error {
		return nil
	})

	assert.NoError(t, err)
}
