package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Budi721/alterra-agmc/v2/lib/database"
	"github.com/Budi721/alterra-agmc/v2/models"
	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()

	if err != nil {
		switch err.Error() {
		case "not found":
			return echo.NewHTTPError(
				http.StatusNotFound,
				models.Response{
					Status: "not found",
					Code:   http.StatusNotFound,
					Data:   []models.Book{},
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
		Data:   books,
	})
}

func GetBookController(c echo.Context) error {
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
	book, err := database.GetBook(uint(convertedId))
	if err != nil {
		switch err.Error() {
		case "not found":
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
		Data:   book,
	})
}

func PostBookController(c echo.Context) error {
	var book models.Book

	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.Response{
				Status: "bad request",
				Code:   http.StatusBadRequest,
				Data:   nil,
			},
		)
	}

	created, err := database.CreateBook(book.ID, book.Title, book.Author, book.Price)
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

func PutBookController(c echo.Context) error {
	var book models.Book
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

	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.Response{
				Status: "bad request",
				Code:   http.StatusBadRequest,
				Data:   nil,
			},
		)
	}

	created, err := database.UpdateBook(uint(convertedId), book.Title, book.Author, book.Price)
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

func DeleteBookController(c echo.Context) error {
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

	book, err := database.DeleteBook(uint(convertedId))
	if err != nil {
		switch err.Error() {
		case "not found":
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
		Data:   book,
	})
}
