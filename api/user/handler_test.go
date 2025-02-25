package user

import (
	"LibraryApi/internal/app/User"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func setup() (*User.MockService, *gin.Engine) {
	router := gin.Default()
	repo := User.GetMockService()
	router.POST("/users", func(c *gin.Context) { Create(c, repo) })
	router.GET("/users/:id", func(c *gin.Context) { GetInfo(c, repo) })
	router.PUT("/users/:id", func(c *gin.Context) { UpdateInfo(c, repo) })
	router.DELETE("/users/:id", func(c *gin.Context) { DeleteUser(c, repo) })
	return repo, router
}

func TestCreate(t *testing.T) {
	repo, router := setup()
	newUser := User.User{
		Name:     "test",
		Email:    "test@test.com",
		Password: "test",
		Role:     "admin",
	}
	requestBody, _ := json.Marshal(newUser)
	//performing request
	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	//assert request
	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response["id"])
	createdUserID := uint64(response["id"].(float64))

	//assert values
	assert.Equal(t, newUser.Name, repo.Users[createdUserID].Name)
	assert.Equal(t, newUser.Email, repo.Users[createdUserID].Email)
}

func TestGetInfo(t *testing.T) {
	repo, router := setup()
	newUser := User.User{
		Name:     "test",
		Email:    "test@test.com",
		Password: "test",
		Role:     "admin",
	}
	id, err := repo.Add(&newUser)
	assert.NoError(t, err)
	newUser.ID = id
	request, _ := http.NewRequest("GET", "/users/"+strconv.FormatUint(id, 10), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	assert.Equal(t, 202, w.Code)
	assert.Equal(t, repo.Users[id], &newUser)
}

func TestUpdateInfo(t *testing.T) {
	repo, router := setup()
	//creating a book manually
	newUser := User.User{
		Name:     "test",
		Email:    "test@test.com",
		Password: "test",
		Role:     "admin",
	}
	id, err := repo.Add(&newUser)
	assert.NoError(t, err)
	newUser.ID = id
	//updated info
	updatedUser := User.User{
		Name:     "testUpdated",
		Email:    "testUpdated@test.com",
		Password: "test",
		Role:     "user",
	}

	requestBody, _ := json.Marshal(&updatedUser)
	request, _ := http.NewRequest("PUT", "/users/"+strconv.FormatUint(id, 10), bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response["id"])

	// Assert book is added to the mock service
	createdUserID := uint64(response["id"].(float64)) // Assert type conversion
	assert.Equal(t, updatedUser.Name, repo.Users[createdUserID].Name)
	assert.Equal(t, updatedUser.Email, repo.Users[createdUserID].Email)
}

func TestDeleteUser(t *testing.T) {
	repo, router := setup()
	newUser := User.User{
		Name:     "test",
		Email:    "test@test.com",
		Password: "test",
		Role:     "admin",
	}
	id, err := repo.Add(&newUser)
	assert.NoError(t, err)
	newUser.ID = id

	request, _ := http.NewRequest("DELETE", "/users/"+strconv.FormatUint(id, 10), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	assert.Equal(t, 204, w.Code)
	_, ok := repo.Users[id]
	assert.False(t, ok)
}
