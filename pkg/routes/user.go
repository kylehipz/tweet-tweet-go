package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProfileHandler(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "hello world!",
		})
}

func RegisterUserRoutes(e *echo.Echo) {
	grp := e.Group("/users")

	grp.GET("/profile", GetProfileHandler)
}
