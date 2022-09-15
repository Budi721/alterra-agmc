package controllers

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"github.com/Budi721/alterra-agmc/v2/lib/database"
	"github.com/Budi721/alterra-agmc/v2/models"
	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	user := models.UserLogin{}
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

	// validator request middleware
	err = c.Validate(user)
	if err != nil {
		return err
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
				Status: "bad request: failed to bind request",
				Code:   http.StatusBadRequest,
				Data:   nil,
			},
		)
	}

	// validator request middleware
	err := c.Validate(user)
	if err != nil {
		return err
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
	// bind payload into model user
	var user models.UserUpdate
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

	// validasi apakah user sesuai dengan yang sedang login
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	if jti, ok := claims["jti"]; ok && jti != id {
		return echo.NewHTTPError(
			http.StatusUnauthorized,
			models.Response{
				Status: "unauthorized",
				Code:   http.StatusUnauthorized,
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

	// validator request middleware
	err = c.Validate(user)
	if err != nil {
		return err
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

	// validasi apakah user sesuai dengan yang sedang login
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	if jti, ok := claims["jti"]; ok && jti != id {
		return echo.NewHTTPError(
			http.StatusUnauthorized,
			models.Response{
				Status: "unauthorized",
				Code:   http.StatusUnauthorized,
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
