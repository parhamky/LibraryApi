package user

import (
	"LibraryApi/internal/app/User"
	"github.com/gin-gonic/gin"
)

type IDParam struct {
	ID uint64 `uri:"id"`
}

func AddRoutes(router *gin.RouterGroup) {
	group := router.Group("/users") // Use "/users" for consistency
	repo := &User.MockService{Users: nil}
	group.POST("/", func(c *gin.Context) { Create(c, repo) })
	group.GET("/:id", func(c *gin.Context) { GetInfo(c, repo) })
	group.PUT("/:id", func(c *gin.Context) { UpdateInfo(c, repo) })
	group.DELETE("/:id", func(c *gin.Context) { DeleteUser(c, repo) })
}
