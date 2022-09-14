package controllers

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"github.com/Budi721/alterra-agmc/v2/lib/database"
	"github.com/Budi721/alterra-agmc/v2/models"
	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.Response{
				Status: "bad request",
				Code:   http.StatusBadRequest,
				Data:   nil,
			},
		)
	}

	token, err := database.LoginUser(user.Email, user.Password)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return echo.NewHTTPError(
				http.StatusNotFound,
				models.Response{
					Status: "not found",
					Code:   http.StatusNotFound,
					Data:   nil,
				},
			)
		default:
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				models.Response{
					Status: "internal server error",
					Code:   http.StatusInternalServerError,
					Data:   nil,
				},
			)
		}
	}

	return c.JSON(http.StatusOK, models.Response{
		Status: "success",
		Code:   http.StatusOK,
		Data: map[string]any{
			"token": token,
		},
	})
}

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return echo.NewHTTPError(
				http.StatusNotFound,
				models.Response{
					Status: "not found",
					Code:   http.StatusNotFound,
					Data:   []models.User{},
				},
			)
		default:
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				models.Response{
					Status: "internal server error",
					Code:   http.StatusInternalServerError,
					Data:   nil,
				},
			)
		}
	}

	return c.JSON(http.StatusOK, models.Response{
		Status: "success",
		Code:   http.StatusOK,
		Data:   users,
	})
}

func GetUserController(c echo.Context) error {
	id := c.Param("id")

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.Response{
				Status: fmt.Sprintf("bad request: %s", err.Error()),
				Code:   http.StatusBadRequest,
				Data:   nil,
			},
		)
	}
	user, err := database.GetUser(uint(convertedId))
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return echo.NewHTTPError(
				http.StatusNotFound,
				models.Response{
					Status: "not found",
					Code:   http.StatusNotFound,
					Data:   nil,
				},
			)
		default:
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				models.Response{
					Status: "internal server error",
					Code:   http.StatusInternalServerError,
					Data:   nil,
				},
			)
		}
	}

	return c.JSON(http.StatusOK, models.Response{
		Status: "success",
		Code:   http.StatusOK,
		Data:   user,
	})
}

func PostUserController(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.Response{
				Status: "bad request",
				Code:   http.StatusBadRequest,
				Data:   nil,
			},
		)
	}

	created, err := database.CreateUser(user.Name, user.Email, user.Password)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			models.Response{
				Status: "internal server error",
				Code:   http.StatusInternalServerError,
				Data:   nil,
			},
		)
	}

	return c.JSON(http.StatusCreated, models.Response{
		Status: "success",
		Code:   http.StatusOK,
		Data:   created,
	})
}

func PutUserController(c echo.Context) error {
	var user models.User
	id := c.Param("id")
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.Response{
				Status: "bad request",
				Code:   http.StatusBadRequest,
				Data:   nil,
			},
		)
	}

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.Response{
				Status: "bad request",
				Code:   http.StatusBadRequest,
				Data:   nil,
			},
		)
	}

	created, err := database.UpdateUser(uint(convertedId), user.Name, user.Email, user.Password)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			models.Response{
				Status: "internal server error",
				Code:   http.StatusInternalServerError,
				Data:   nil,
			},
		)
	}

	return c.JSON(http.StatusOK, models.Response{
		Status: "success",
		Code:   http.StatusOK,
		Data:   created,
	})
}

func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.Response{
				Status: "bad request",
				Code:   http.StatusBadRequest,
				Data:   nil,
			},
		)
	}

	user, err := database.DeleteUser(uint(convertedId))
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return echo.NewHTTPError(
				http.StatusNotFound,
				models.Response{
					Status: "not found",
					Code:   http.StatusNotFound,
					Data:   nil,
				},
			)
		default:
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				models.Response{
					Status: "internal server error",
					Code:   http.StatusInternalServerError,
					Data:   nil,
				},
			)
		}
	}

	return c.JSON(http.StatusOK, models.Response{
		Status: "success",
		Code:   http.StatusOK,
		Data:   user,
	})
}
