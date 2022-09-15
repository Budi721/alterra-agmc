package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Budi721/alterra-agmc/v2/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.Response{
				Status: fmt.Sprintf("bad request: %s", err.Error()),
				Code:   http.StatusBadRequest,
				Data:   nil,
			},
		)
	}

	return nil
}
