package user

import (
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("/login", h.LoginUserController)
	g.GET("", h.GetUsersController)
	g.GET("/:id", h.GetUserController)
	g.POST("", h.PostUserController)
	g.PUT("/:id", h.PutUserController)
	g.DELETE("/:id", h.DeleteUserController)
}
