package user

import (
	"github.com/Budi721/alterra-agmc/v6/internal/dto"
	"github.com/Budi721/alterra-agmc/v6/internal/middleware"
	"github.com/Budi721/alterra-agmc/v6/internal/pkg/util"
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("/login", h.LoginUserController)
	g.GET("", h.GetUsersController, middleware.JWTMiddleware(dto.JWTClaims{}, util.JwtSecret))
	g.GET("/:id", h.GetUserController, middleware.JWTMiddleware(dto.JWTClaims{}, util.JwtSecret))
	g.POST("", h.PostUserController)
	g.PUT("/:id", h.PutUserController, middleware.JWTMiddleware(dto.JWTClaims{}, util.JwtSecret))
	g.DELETE("/:id", h.DeleteUserController, middleware.JWTMiddleware(dto.JWTClaims{}, util.JwtSecret))
}
