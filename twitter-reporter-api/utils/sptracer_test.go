package utils

import (
	"context"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSpanTracer(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/pedidos/v1/arquivo-p/:codigoFilial/receber")
	c.SetParamNames("codigoFilial")
	c.SetParamValues("111")

	err := SpanTracer(c, "", func(cx context.Context, span tracer.Span) error {
		return nil
	})

	assert.NoError(t, err)
}
