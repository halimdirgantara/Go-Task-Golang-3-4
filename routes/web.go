package routes

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"task-list/controllers"
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

//route function
func Web() {
	e := echo.New()
	e.Static("/","assets")
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.Renderer = renderer

	e.GET("/", Controllers.Index)
	e.GET("/tasks", Controllers.TaskList)

	// e.GET("/tasks", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Task Page")
	// })
	
	e.GET("/task/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "Task id: ")
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))

}
