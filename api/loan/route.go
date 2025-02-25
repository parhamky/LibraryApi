package loan

import (
	"LibraryApi/internal/app/Loan"
	"github.com/gin-gonic/gin"
)

type IDParam struct {
	ID uint64 `uri:"id"`
}

func AddRoutes(router *gin.RouterGroup) {
	group := router.Group("/loans") // Use "/loans" for consistency
	repo := &Loan.MockService{Loans: nil}
	group.POST("/", func(c *gin.Context) { Create(c, repo) })
	group.GET("/:id", func(c *gin.Context) { GetInfo(c, repo) })
	group.PUT("/:id", func(c *gin.Context) { UpdateInfo(c, repo) })
	group.DELETE("/:id", func(c *gin.Context) { DeleteLoan(c, repo) })
}
