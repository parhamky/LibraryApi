package loan

import (
	"LibraryApi/internal/app/Loan"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getService() Loan.Service {
	return Loan.GetService()
}

func Create(c *gin.Context, loan *Loan.MockService) {
	var service = getService()
	if loan.Loans != nil {
		service = loan
	}
	var payload Loan.Loan

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

func GetInfo(c *gin.Context, loan *Loan.MockService) {
	var service = getService()
	if loan.Loans != nil {
		service = loan
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

func UpdateInfo(c *gin.Context, loan *Loan.MockService) {
	var service = getService()
	if loan.Loans != nil {
		service = loan
	}
	var payload Loan.Loan
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

func DeleteLoan(c *gin.Context, loan *Loan.MockService) {
	var service = getService()
	if loan.Loans != nil {
		service = loan
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
