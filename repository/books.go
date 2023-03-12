package repository

import (
	"github.com/libraryAPI/entity"
	"sort"
)

type Books interface {
	GetBooks() ([]entity.Book, error)
	SetBooks(newLibrary []entity.Book) error
}

type books struct {
	books []entity.Book
}

func (u *books) GetBooks() ([]entity.Book, error) {
	sort.Sort(u)
	return u.books, nil
}

func (u *books) SetBooks(newLibrary []entity.Book) error {
	u.books = newLibrary
	return nil
}

func NewLibrary() Books {
	return &books{books: []entity.Book{
		{
			Name:   "Rage",
			Author: "Stephen King",
			Year:   1977,
		},
		{
			Name:   "Philosopher's Stone",
			Author: "J. K. Rowling",
			Year:   1997,
		},
		{
			Name:   "All Quiet on the Western Front",
			Author: "Erich Maria Remarque",
			Year:   1929,
		},
	}}
}

func (l books) Len() int {
	return len(l.books)
}
func (l books) Less(i, j int) bool {
	iYear := l.books[i].Year
	jYear := l.books[j].Year
	return iYear < jYear
}
func (l books) Swap(i, j int) {
	l.books[i], l.books[j] = l.books[j], l.books[i]
}
