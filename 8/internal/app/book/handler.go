package book

import (
	"github.com/Budi721/alterra-agmc/v6/internal/factory"
	"github.com/Budi721/alterra-agmc/v6/internal/model"
	res "github.com/Budi721/alterra-agmc/v6/pkg/util/response"
	"github.com/labstack/echo/v4"
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

func (b handler) GetBooksController(c echo.Context) error {
	books, err := b.service.GetBooks()

	if err != nil {
		switch err.Error() {
		case "not found":
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

	return c.JSON(http.StatusOK, res.SuccessResponse(books))
}

func (b handler) GetBookController(c echo.Context) error {
	id := c.Param("id")

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			res.ErrorResponse(err),
		)
	}
	book, err := b.service.GetBook(uint(convertedId))
	if err != nil {
		switch err.Error() {
		case "not found":
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

	return c.JSON(http.StatusOK, res.SuccessResponse(book))
}

func (b handler) PostBookController(c echo.Context) error {
	var book model.Book

	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			res.ErrorResponse(err),
		)
	}

	// validator request middleware
	if err := c.Validate(book); err != nil {
		return err
	}

	created, err := b.service.CreateBook(book.ID, book.Title, book.Author, book.Price)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			res.ErrorResponse(err),
		)
	}

	return c.JSON(http.StatusCreated, res.SuccessResponse(created))
}

func (b handler) PutBookController(c echo.Context) error {
	var book model.Book
	id := c.Param("id")
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			res.ErrorResponse(err),
		)
	}

	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			res.ErrorResponse(err),
		)
	}

	// validator request middleware
	if err := c.Validate(book); err != nil {
		return err
	}

	created, err := b.service.UpdateBook(uint(convertedId), book.Title, book.Author, book.Price)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			res.ErrorResponse(err),
		)
	}

	return c.JSON(http.StatusOK, res.SuccessResponse(created))
}

func (b handler) DeleteBookController(c echo.Context) error {
	id := c.Param("id")
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			res.ErrorResponse(err),
		)
	}

	book, err := b.service.DeleteBook(uint(convertedId))
	if err != nil {
		switch err.Error() {
		case "not found":
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

	return c.JSON(http.StatusOK, res.SuccessResponse(book))
}
