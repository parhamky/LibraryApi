package book

import (
	"LibraryApi/internal/app/Book"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func setup() (*Book.MockService, *gin.Engine) {
	router := gin.Default()
	repo := Book.GetMockService()
	router.POST("/books", func(c *gin.Context) { Create(c, repo) })
	router.GET("/books/:id", func(c *gin.Context) { GetInfo(c, repo) })
	router.PUT("/books/:id", func(c *gin.Context) { UpdateInfo(c, repo) })
	router.DELETE("/books/:id", func(c *gin.Context) { DeleteBook(c, repo) })
	return repo, router
}

func TestCreate(t *testing.T) {
	repo, router := setup()
	newBook := Book.Book{
		Title:       "Test Book",
		Author:      "Test Author",
		ISBN:        "978-3-16-148410-0",
		IsAvailable: true,
	}
	requestBody, _ := json.Marshal(newBook)
	//performing request
	request, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	//assert request
	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response["id"])
	createdBookID := uint64(response["id"].(float64))

	//assert values
	assert.Equal(t, newBook.Title, repo.Books[createdBookID].Title)
	assert.Equal(t, newBook.Author, repo.Books[createdBookID].Author)
	assert.Equal(t, newBook.ISBN, repo.Books[createdBookID].ISBN)
	assert.Equal(t, newBook.IsAvailable, repo.Books[createdBookID].IsAvailable)
}

func TestGetInfo(t *testing.T) {
	repo, router := setup()
	newBook := Book.Book{
		Title:       "Test Book 2",
		Author:      "Test Author 2",
		ISBN:        "972-3-26-148412-0",
		IsAvailable: true,
	}
	id, err := repo.Add(&newBook)
	assert.NoError(t, err)
	newBook.Id = id
	request, _ := http.NewRequest("GET", "/books/"+strconv.FormatUint(id, 10), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	assert.Equal(t, 202, w.Code)
	assert.Equal(t, repo.Books[id], &newBook)
}

func TestUpdateInfo(t *testing.T) {
	repo, router := setup()
	//creating a book manually
	newBook := Book.Book{
		Title:       "Test Book 3",
		Author:      "Test Author 3",
		ISBN:        "978-4-13-148430-3",
		IsAvailable: true,
	}
	id, err := repo.Add(&newBook)
	assert.NoError(t, err)
	newBook.Id = id
	//updated info
	updatedBook := Book.Book{
		Title:       "Test Book 4",
		Author:      "Test Author 4",
		ISBN:        "978-3-13-348330-3",
		IsAvailable: true,
	}

	requestBody, _ := json.Marshal(&updatedBook)
	request, _ := http.NewRequest("PUT", "/books/"+strconv.FormatUint(id, 10), bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	log.Println("++++++++++++++++++++++++++log mock+++++++++++++++++++++++++++++++")
	log.Println(response["id"])
	log.Println("++++++++++++++++++++++++++end+++++++++++++++++++++++++++++++")
	assert.NoError(t, err)
	assert.NotNil(t, response["id"])

	// Assert book is added to the mock service
	createdBookID := uint64(response["id"].(float64)) // Assert type conversion
	assert.Equal(t, updatedBook.Title, repo.Books[createdBookID].Title)
	assert.Equal(t, updatedBook.Author, repo.Books[createdBookID].Author)
	assert.Equal(t, updatedBook.ISBN, repo.Books[createdBookID].ISBN)
	assert.Equal(t, updatedBook.IsAvailable, repo.Books[createdBookID].IsAvailable)
}

func TestDeleteBook(t *testing.T) {
	repo, router := setup()
	newBook := Book.Book{
		Title:       "Test Book 5",
		Author:      "Test Author 5",
		ISBN:        "978-3-13-348330-3",
		IsAvailable: true,
	}
	id, err := repo.Add(&newBook)
	assert.NoError(t, err)
	newBook.Id = id

	request, _ := http.NewRequest("DELETE", "/books/"+strconv.FormatUint(id, 10), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	assert.Equal(t, 204, w.Code)
	_, ok := repo.Books[id]
	assert.False(t, ok)
}
