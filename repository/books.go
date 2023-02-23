package repository

import "github.com/libraryAPI/entity"

type Library interface {
	GetBooks() ([]entity.Book, error)
	SetBooks(newLibrary []entity.Book) error
}

type library struct {
	books []entity.Book
}

func (u *library) GetBooks() ([]entity.Book, error) {
	return u.books, nil
}

func (u *library) SetBooks(newLibrary []entity.Book) error {
	u.books = newLibrary
	return nil
}

func NewLibrary() Library {
	return &library{books: []entity.Book{
		{
			Name:   "Rage",
			Author: "Stephen King",
			Year:   "1977",
		},
		{
			Name:   "Philosopher's Stone",
			Author: "J. K. Rowling",
			Year:   "1997",
		},
		{
			Name:   "All Quiet on the Western Front",
			Author: "Erich Maria Remarque",
			Year:   "1929",
		},
	}}
}
