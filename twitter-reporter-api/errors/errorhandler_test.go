package errors

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCriticaHttpErrorHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	CriticaHttpErrorHandler()(NewError(http.StatusInternalServerError, "errorr", nil), c)
	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
}

func TestCriticaHttpErrorHandler2(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	CriticaHttpErrorHandler()(echo.NewHTTPError(http.StatusInternalServerError, "failure"), c)
	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
}