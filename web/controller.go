package web

import (
	"github.com/libraryAPI/entity"
	"github.com/libraryAPI/repository"
)

type Controller struct {
	bookLibrary repository.Library
}

func (c *Controller) GetBooks() ([]entity.Book, error) {
	return c.bookLibrary.GetBooks()
}
