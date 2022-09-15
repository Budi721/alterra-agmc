package routes

import (
	"github.com/Budi721/alterra-agmc/v2/controllers"
	"github.com/Budi721/alterra-agmc/v2/middlewares"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	// Add middleware
	e.Pre(middleware.RemoveTrailingSlash())
	middlewares.LogMiddleware(e)
	e.Validator = &middlewares.CustomValidator{Validator: validator.New()}
	// Grouping route api version 1
	v1 := e.Group("/v1")
	v1.POST("/login", controllers.LoginUserController)
	// Grouping route for user
	userGroup := v1.Group("/users")
	userGroup.GET("", controllers.GetUsersController, middlewares.JWTMiddleware)
	userGroup.GET("/:id", controllers.GetUserController, middlewares.JWTMiddleware)
	userGroup.POST("", controllers.PostUserController)
	userGroup.PUT("/:id", controllers.PutUserController, middlewares.JWTMiddleware)
	userGroup.DELETE("/:id", controllers.DeleteUserController, middlewares.JWTMiddleware)
	// Grouping route for book
	bookGroup := v1.Group("/books")
	bookGroup.GET("", controllers.GetBooksController)
	bookGroup.GET("/:id", controllers.GetBookController)
	bookGroup.POST("", controllers.PostBookController, middlewares.JWTMiddleware)
	bookGroup.PUT("/:id", controllers.PutBookController, middlewares.JWTMiddleware)
	bookGroup.DELETE("/:id", controllers.DeleteBookController, middlewares.JWTMiddleware)

	return e
}
