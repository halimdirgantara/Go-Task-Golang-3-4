package Controllers

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func TaskList(c echo.Context) error {
	e := echo.New()
	e.Renderer = &TemplateRenderer{templates: template.Must(template.ParseFiles(
		"views/tasks.html",
		"views/home.html",
		"views/header.html",
		"views/footer.html",
	)),
	}
	return c.Render(http.StatusOK, "tasks.html", map[string]interface{}{
	})
}

func TaskDetail(c echo.Context) error {
	e := echo.New()
	e.Renderer = &TemplateRenderer{templates: template.Must(template.ParseFiles(
		"views/tasks.html",
		"views/home.html",
		"views/header.html",
		"views/footer.html",
	)),
	}
	return c.Render(http.StatusOK, "task-detail.html", map[string]interface{}{
	})
}