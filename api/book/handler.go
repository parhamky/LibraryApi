package book

import (
	"LibraryApi/internal/app/Book"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getService() Book.Service {
	return Book.GetService()
}

func Create(c *gin.Context, book *Book.MockService) {
	var service = getService()
	if book.Books != nil {
		service = book
	}
	var payload Book.Book

	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nil body not expected"})
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := service.Add(&payload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": float64(id)})
}

func GetInfo(c *gin.Context, book *Book.MockService) {
	service := getService()
	if book.Books != nil {
		service = book
	}
	var payload IDParam

	if err := c.ShouldBindUri(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := service.Read(&payload.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(202, res)
}

func UpdateInfo(c *gin.Context, book *Book.MockService) {
	service := getService()
	if book.Books != nil {
		service = book
	}
	var payload Book.Book
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var idparam IDParam
	if err := c.ShouldBindUri(&idparam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := service.Update(&payload, &idparam.ID); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusAccepted, gin.H{"id": float64(idparam.ID)})
}

func DeleteBook(c *gin.Context, book *Book.MockService) {
	service := getService()
	if book.Books != nil {
		service = book
	}
	var payload IDParam

	if err := c.ShouldBindUri(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := service.Delete(&payload.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	res := fmt.Sprintf("id:%d deleted", &payload)

	c.JSON(204, &res)
}
