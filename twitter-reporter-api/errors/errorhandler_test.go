package errors

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpErrorHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	HttpErrorHandler()(NewError(http.StatusInternalServerError, "error", nil), c)
	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
}

func TestHttpErrorHandler2(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	HttpErrorHandler()(NewError(http.StatusInternalServerError, "failure", nil), c)
	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
}
