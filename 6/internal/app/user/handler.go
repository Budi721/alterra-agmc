package user

import (
	"errors"

	"github.com/Budi721/alterra-agmc/v6/internal/dto"
	"github.com/Budi721/alterra-agmc/v6/internal/factory"
	res "github.com/Budi721/alterra-agmc/v6/pkg/util/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h handler) LoginUserController(c echo.Context) error {
	user := dto.UserLogin{}
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			res.SuccessResponse(user),
		)
	}

	// validator request middleware
	err = c.Validate(user)
	if err != nil {
		return err
	}

	token, err := h.service.LoginByEmailAndPassword(&dto.UserLogin{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return echo.NewHTTPError(
				http.StatusNotFound,
				res.ErrorResponse(err),
			)
		default:
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				res.ErrorResponse(err),
			)
		}
	}

	return c.JSON(http.StatusOK, res.SuccessResponse(map[string]any{
		"token": token,
	}))
}

func (h handler) GetUsersController(c echo.Context) error {
	users, err := h.service.GetUsers()

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return echo.NewHTTPError(
				http.StatusNotFound,
				res.ErrorResponse(err),
			)
		default:
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				res.ErrorResponse(err),
			)
		}
	}

	return c.JSON(http.StatusOK, res.SuccessResponse(users))
}

func (h handler) GetUserController(c echo.Context) error {
	id := c.Param("id")

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			res.ErrorResponse(err),
		)
	}
	user, err := h.service.GetUser(uint(convertedId))
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return echo.NewHTTPError(
				http.StatusNotFound,
				res.ErrorResponse(err),
			)
		default:
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				res.ErrorResponse(err),
			)
		}
	}

	return c.JSON(http.StatusOK, res.SuccessResponse(user))
}

func (h handler) PostUserController(c echo.Context) error {
	var user dto.User

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			res.ErrorResponse(err),
		)
	}

	// validator request middleware
	err := c.Validate(user)
	if err != nil {
		return err
	}

	created, err := h.service.CreateUser(&user)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			res.ErrorResponse(err),
		)
	}

	return c.JSON(http.StatusCreated, res.SuccessResponse(created))
}

func (h handler) PutUserController(c echo.Context) error {
	// bind payload into model user
	var user dto.UserUpdate
	id := c.Param("id")
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			res.ErrorResponse(err),
		)
	}

	// validasi apakah user sesuai dengan yang sedang login
	//token := c.Get("user").(*jwt.Token)
	//claims := token.Claims.(jwt.MapClaims)
	//if jti, ok := claims["jti"]; ok && jti != id {
	//	return echo.NewHTTPError(
	//		http.StatusUnauthorized,
	//		res.ErrorResponse(err),
	//	)
	//}

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			res.ErrorResponse(err),
		)
	}

	// validator request middleware
	err = c.Validate(user)
	if err != nil {
		return err
	}

	created, err := h.service.UpdateUser(uint(convertedId), &dto.UserUpdate{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			res.ErrorResponse(err),
		)
	}

	return c.JSON(http.StatusOK, res.SuccessResponse(created))
}

func (h handler) DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			res.ErrorResponse(err),
		)
	}

	// validasi apakah user sesuai dengan yang sedang login
	//token := c.Get("user").(*jwt.Token)
	//claims := token.Claims.(jwt.MapClaims)
	//
	//if jti, ok := claims["jti"]; ok && jti != id {
	//	return echo.NewHTTPError(
	//		http.StatusUnauthorized,
	//		res.ErrorResponse(err),
	//	)
	//}

	user, err := h.service.DeleteUser(uint(convertedId))
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return echo.NewHTTPError(
				http.StatusNotFound,
				res.ErrorResponse(err),
			)
		default:
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				res.ErrorResponse(err),
			)
		}
	}

	return c.JSON(http.StatusOK, res.SuccessResponse(user))
}
