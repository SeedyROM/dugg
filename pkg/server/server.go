package server

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer() *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func Start(host string, port uint16) {
	e := echo.New()

	e.Static("/", "public/")

	renderer := NewTemplateRenderer()
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"Title": "Dugg is here! I am gay!!! Have no fear!!!",
		})
	})
	hostAddress := fmt.Sprintf("%s:%d", host, port)
	e.Logger.Fatal(e.Start(hostAddress))
}
