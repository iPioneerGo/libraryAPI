package main

import (
	"github.com/gin-gonic/gin"
	"github.com/libraryAPI/entity"
	"github.com/libraryAPI/repository"
	"log"
	"net/http"
)

var books repository.Books

func getBooks(c *gin.Context) {
	bookList, err := books.GetBooks()
	if err != nil {
		log.Panic(err)
	}
	c.IndentedJSON(http.StatusOK, bookList)
}

func setBooks(c *gin.Context) {
	var bookList []entity.Book
	if err := c.BindJSON(&bookList); err != nil {
		log.Panic(err)
	}
	books.SetBooks(bookList)
	c.IndentedJSON(http.StatusOK, "Books were set")
}

func main() {
	books = repository.NewLibrary()
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", setBooks)
	router.Run("localhost:8080")
}
