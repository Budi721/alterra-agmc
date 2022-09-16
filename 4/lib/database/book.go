package database

import (
	"fmt"

	"github.com/Budi721/alterra-agmc/v2/lib/mock"
	"github.com/Budi721/alterra-agmc/v2/models"
)

func GetBooks() ([]models.Book, error) {
	// check if static data available
	if len(mock.Books) == 0 {
		return []models.Book{}, fmt.Errorf("not found")
	}

	return mock.Books, nil
}

func GetBook(id uint) (models.Book, error) {
	// find book based on id
	for _, b := range mock.Books {
		if b.ID == id {
			return b, nil
		}
	}

	// check if available book id
	return models.Book{}, fmt.Errorf("not found")
}

func CreateBook(id uint, title string, author string, price uint) (models.Book, error) {
	// create new instance book from request
	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Price:  price,
	}
	// appending mocking data
	mock.Books = append(mock.Books, book)
	// returning instance of book
	return book, nil
}

func UpdateBook(id uint, title string, author string, price uint) (models.Book, error) {
	for _, b := range mock.Books {
		if b.ID == id {
			book := mock.Books[id]
			mock.Books[id] = models.Book{
				ID:     book.ID,
				Title:  title,
				Author: author,
				Price:  price,
			}
			// return book with same id
			return book, nil
		}
	}

	// check if available book id
	return models.Book{}, fmt.Errorf("not found")
}

func DeleteBook(id uint) (models.Book, error) {
	for i, b := range mock.Books {
		if b.ID == id {
			book := mock.Books[i]
			mock.Books = append(mock.Books[:i], mock.Books[i+1:]...)
			return book, nil
		}
	}

	// check if available book id
	return models.Book{}, fmt.Errorf("not found")
}
