package routes

import (
	"github.com/Budi721/alterra-agmc/v2/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	// Grouping route api version 1
	v1 := e.Group("/v1")
	// Grouping route for user
	userGroup := v1.Group("/users")
	userGroup.GET("", controllers.GetUsersController)
	userGroup.GET("/:id", controllers.GetUserController)
	userGroup.POST("", controllers.PostUserController)
	userGroup.PUT("/:id", controllers.PutUserController)
	userGroup.DELETE("/:id", controllers.DeleteUserController)
	// Grouping route for book
	bookGroup := v1.Group("/books")
	bookGroup.GET("", controllers.GetBooksController)
	bookGroup.GET("/:id", controllers.GetBookController)
	bookGroup.POST("", controllers.PostBookController)
	bookGroup.PUT("/:id", controllers.PutBookController)
	bookGroup.DELETE("/:id", controllers.DeleteBookController)

	return e
}
