package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/libraryAPI/entity"
	"github.com/libraryAPI/repository"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBooksEndpoint(t *testing.T) {
	router := SetUpRouter()
	books = repository.NewLibrary()

	router.GET("/books", getBooks)
	request, _ := http.NewRequest("GET", "/books", nil)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	var bookList []entity.Book
	json.Unmarshal(writer.Body.Bytes(), &bookList)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.NotEmpty(t, bookList)

}

func TestSetBooksEndpoint(t *testing.T) {
	router := SetUpRouter()

	router.POST("/books", setBooks)
	mockResponse := "\"Books were set\""
	books := []entity.Book{
		{
			Name:   "Rage",
			Author: "Stephen King",
			Year:   1900,
		},
		{
			Name:   "Philosopher's Stone",
			Author: "J. K. Rowling",
			Year:   1900,
		},
		{
			Name:   "All Quiet on the Western Front",
			Author: "Erich Maria Remarque",
			Year:   1900,
		},
	}

	jsonValue, _ := json.Marshal(books)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
