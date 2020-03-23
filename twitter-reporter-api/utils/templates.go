package utils

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
)

type templateRenderer struct {
	templates *template.Template
}

func (t *templateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplateRenderer() *templateRenderer {
	return &templateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
}
