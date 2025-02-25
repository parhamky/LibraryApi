package api

import (
	"LibraryApi/api/book"
	"LibraryApi/api/loan"
	"LibraryApi/api/user"
	"LibraryApi/internal/config"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	r := gin.Default()
	r.Use(cors.Default())

	v1 := r.Group("/v1")

	book.AddRoutes(v1)
	loan.AddRoutes(v1)
	user.AddRoutes(v1)

	return &Router{r}
}

func (r *Router) Serve() error {
	httpConfig := config.LoadHttpConfig()

	listenAddr := fmt.Sprintf("%s:%s", httpConfig.Url, httpConfig.Port)

	return r.Run(listenAddr)
}
