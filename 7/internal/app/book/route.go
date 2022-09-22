package book

import (
	"github.com/Budi721/alterra-agmc/v6/internal/dto"
	"github.com/Budi721/alterra-agmc/v6/internal/middleware"
	"github.com/Budi721/alterra-agmc/v6/internal/pkg/util"
	"github.com/labstack/echo/v4"
)

func (b *handler) Route(g *echo.Group) {
	g.GET("", b.GetBooksController)
	g.GET("/:id", b.GetBookController)
	g.POST("", b.PostBookController, middleware.JWTMiddleware(dto.JWTClaims{}, util.JwtSecret))
	g.PUT("/:id", b.PutBookController, middleware.JWTMiddleware(dto.JWTClaims{}, util.JwtSecret))
	g.DELETE("/:id", b.DeleteBookController, middleware.JWTMiddleware(dto.JWTClaims{}, util.JwtSecret))
}
