package factory

import (
	"github.com/Budi721/alterra-agmc/v6/database"
	"github.com/Budi721/alterra-agmc/v6/internal/repository"
)

type Factory struct {
	UserRepository repository.User
	BookRepository repository.Book
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		repository.NewUserRepository(db),
		repository.NewBookRepository(db),
	}
}
