package repository

import (
	"fmt"

	"github.com/Budi721/alterra-agmc/v6/internal/model"
	"gorm.io/gorm"
)

// Books mocking static data for endpoint book
var books = []model.Book{
	{ID: uint(1), Title: "Anak Singkong", Author: "Chairil Tanjung", Price: uint(50000)},
	{ID: uint(2), Title: "Garis Waktu", Author: "Fiersa Besari", Price: uint(35000)},
}

type Book interface {
	GetBooks() ([]model.Book, error)
	GetBook(id uint) (*model.Book, error)
	CreateBook(id uint, title string, author string, price uint) (*model.Book, error)
	UpdateBook(id uint, title string, author string, price uint) (*model.Book, error)
	DeleteBook(id uint) (*model.Book, error)
}

type book struct {
	Db *gorm.DB
}

func (b book) GetBooks() ([]model.Book, error) {
	// check if static data available
	if len(books) == 0 {
		return []model.Book{}, fmt.Errorf("not found")
	}

	return books, nil
}

func (b book) GetBook(id uint) (*model.Book, error) {
	// find book based on id
	for _, b := range books {
		if b.ID == id {
			return &b, nil
		}
	}

	// check if available book id
	return &model.Book{}, fmt.Errorf("not found")
}

func (b book) CreateBook(id uint, title string, author string, price uint) (*model.Book, error) {
	// create new instance book from request
	book := model.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Price:  price,
	}
	// appending mocking data
	books = append(books, book)
	// returning instance of book
	return &book, nil
}

func (b book) UpdateBook(id uint, title string, author string, price uint) (*model.Book, error) {
	for _, b := range books {
		if b.ID == id {
			book := books[id]
			books[id] = model.Book{
				ID:     book.ID,
				Title:  title,
				Author: author,
				Price:  price,
			}
			// return book with same id
			return &books[id], nil
		}
	}

	// check if available book id
	return &model.Book{}, fmt.Errorf("not found")
}

func (b book) DeleteBook(id uint) (*model.Book, error) {
	for i, b := range books {
		if b.ID == id {
			book := books[i]
			books = append(books[:i], books[i+1:]...)
			return &book, nil
		}
	}

	// check if available book id
	return &model.Book{}, fmt.Errorf("not found")
}

func NewBookRepository(db *gorm.DB) Book {
	return &book{Db: db}
}
