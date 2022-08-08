package Controllers

import (

	"net/http"

	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
	})
}