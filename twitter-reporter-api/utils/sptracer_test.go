package utils

import (
	"context"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"go.elastic.co/apm"
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

	err := SpanTracer(c.Request().Context(), "", "", func(ctx context.Context, span *apm.Span) error {
		return nil
	})

	assert.NoError(t, err)
}
