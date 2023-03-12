package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/libraryAPI/entity"
	"github.com/libraryAPI/repository"
	"github.com/stretchr/testify/assert"
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

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
