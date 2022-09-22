package book

import (
	"github.com/Budi721/alterra-agmc/v6/internal/factory"
	"github.com/Budi721/alterra-agmc/v6/internal/model"
	"github.com/Budi721/alterra-agmc/v6/internal/repository"
	res "github.com/Budi721/alterra-agmc/v6/pkg/util/response"
)

type service struct {
	BookRepository repository.Book
}

func (s service) GetBooks() ([]model.Book, error) {
	books, err := s.BookRepository.GetBooks()
	if err != nil {
		return []model.Book{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return books, nil
}

func (s service) GetBook(id uint) (*model.Book, error) {
	book, err := s.BookRepository.GetBook(id)
	if err != nil {
		return &model.Book{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return book, nil
}

func (s service) CreateBook(id uint, title string, author string, price uint) (*model.Book, error) {
	book, err := s.BookRepository.CreateBook(id, title, author, price)
	if err != nil {
		return &model.Book{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return book, nil
}

func (s service) UpdateBook(id uint, title string, author string, price uint) (*model.Book, error) {
	book, err := s.BookRepository.UpdateBook(id, title, author, price)
	if err != nil {
		return &model.Book{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return book, nil
}

func (s service) DeleteBook(id uint) (*model.Book, error) {
	book, err := s.BookRepository.DeleteBook(id)
	if err != nil {
		return &model.Book{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return book, nil
}

type Service interface {
	GetBooks() ([]model.Book, error)
	GetBook(id uint) (*model.Book, error)
	CreateBook(id uint, title string, author string, price uint) (*model.Book, error)
	UpdateBook(id uint, title string, author string, price uint) (*model.Book, error)
	DeleteBook(id uint) (*model.Book, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		BookRepository: f.BookRepository,
	}
}
