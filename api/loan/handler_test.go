package loan

import (
	"LibraryApi/internal/app/Loan"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func setup() (*Loan.MockService, *gin.Engine) {
	router := gin.Default()
	repo := Loan.GetMockService()
	router.POST("/loans", func(c *gin.Context) { Create(c, repo) })
	router.GET("/loans/:id", func(c *gin.Context) { GetInfo(c, repo) })
	router.PUT("/loans/:id", func(c *gin.Context) { UpdateInfo(c, repo) })
	router.DELETE("/loans/:id", func(c *gin.Context) { DeleteLoan(c, repo) })
	return repo, router
}

func TestCreate(t *testing.T) {
	repo, router := setup()
	newLoan := Loan.Loan{
		UserID:     1,
		BookID:     2,
		LoanedAt:   time.Now(),
		DueDate:    time.Now().Add(24 * time.Hour),
		ReturnDate: time.Time{},
	}
	requestBody, _ := json.Marshal(newLoan)
	//performing request
	request, _ := http.NewRequest("POST", "/loans", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	//assert request
	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response["id"])
	createdLoanID := uint64(response["id"].(float64))

	//assert values
	assert.Equal(t, newLoan.UserID, repo.Loans[createdLoanID].UserID)
	assert.Equal(t, newLoan.BookID, repo.Loans[createdLoanID].BookID)
}

func TestGetInfo(t *testing.T) {
	repo, router := setup()
	newLoan := Loan.Loan{
		UserID:     1,
		BookID:     2,
		LoanedAt:   time.Now(),
		DueDate:    time.Now().Add(24 * time.Hour),
		ReturnDate: time.Time{},
	}
	id, err := repo.Add(&newLoan)
	assert.NoError(t, err)
	newLoan.ID = id
	request, _ := http.NewRequest("GET", "/loans/"+strconv.FormatUint(id, 10), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	assert.Equal(t, 202, w.Code)
	assert.Equal(t, repo.Loans[id], &newLoan)
}

func TestUpdateInfo(t *testing.T) {
	repo, router := setup()
	//creating a book manually
	newLoan := Loan.Loan{
		UserID:     1,
		BookID:     2,
		LoanedAt:   time.Now(),
		DueDate:    time.Now().Add(24 * time.Hour),
		ReturnDate: time.Time{},
	}
	id, err := repo.Add(&newLoan)
	assert.NoError(t, err)
	newLoan.ID = id
	//updated info
	updatedLoan := Loan.Loan{
		UserID:     1,
		BookID:     3,
		LoanedAt:   time.Now(),
		DueDate:    time.Now().Add(48 * time.Hour),
		ReturnDate: time.Time{},
	}

	requestBody, _ := json.Marshal(&updatedLoan)
	request, _ := http.NewRequest("PUT", "/loans/"+strconv.FormatUint(id, 10), bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response["id"])

	// Assert book is added to the mock service
	createdLoanID := uint64(response["id"].(float64)) // Assert type conversion
	assert.Equal(t, updatedLoan.UserID, repo.Loans[createdLoanID].UserID)
	assert.Equal(t, updatedLoan.BookID, repo.Loans[createdLoanID].BookID)
}

func TestDeleteLoan(t *testing.T) {
	repo, router := setup()
	newLoan := Loan.Loan{
		UserID:     1,
		BookID:     2,
		LoanedAt:   time.Now(),
		DueDate:    time.Now().Add(24 * time.Hour),
		ReturnDate: time.Time{},
	}
	id, err := repo.Add(&newLoan)
	assert.NoError(t, err)
	newLoan.ID = id

	request, _ := http.NewRequest("DELETE", "/loans/"+strconv.FormatUint(id, 10), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	assert.Equal(t, 204, w.Code)
	_, ok := repo.Loans[id]
	assert.False(t, ok)
}
