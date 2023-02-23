package main

import (
	"github.com/gin-gonic/gin"
	"github.com/libraryAPI/repository"
	"net/http"
)

var books repository.Library

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	books = repository.NewLibrary()
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}
