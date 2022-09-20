package http

import (
	"github.com/Budi721/alterra-agmc/v6/internal/app/book"
	"github.com/Budi721/alterra-agmc/v6/internal/app/user"
	"github.com/Budi721/alterra-agmc/v6/internal/factory"
	"github.com/Budi721/alterra-agmc/v6/pkg/util"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})
	v1 := e.Group("/api/v1")
	user.NewHandler(f).Route(v1.Group("/users"))
	book.NewHandler(f).Route(v1.Group("/books"))
}
