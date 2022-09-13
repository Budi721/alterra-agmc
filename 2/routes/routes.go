package routes

import (
	"github.com/Budi721/alterra-agmc/v2/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUserControllers)

	return e
}
