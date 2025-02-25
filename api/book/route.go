package book

import (
	"LibraryApi/internal/app/Book"
	"github.com/gin-gonic/gin"
)

type IDParam struct {
	ID uint64 `uri:"id"`
}

func AddRoutes(router *gin.RouterGroup) {
	group := router.Group("/books")
	repo := &Book.MockService{Books: nil}
	group.POST("/", func(c *gin.Context) { Create(c, repo) })
	group.GET("/:id", func(c *gin.Context) { GetInfo(c, repo) })
	group.PUT("/:id", func(c *gin.Context) { UpdateInfo(c, repo) })
	group.DELETE("/:id", func(c *gin.Context) { DeleteBook(c, repo) })
}
