package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mogfee/blog-server/internal/middleware"
	"github.com/mogfee/blog-server/internal/middleware/exception"
	"github.com/mogfee/blog-server/internal/routers/api"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	//r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Tracing())
	r.Use(middleware.AccessLog())
	r.Use(exception.SetUp())
	tag := api.NewTagController()
	article := api.NewArticle()
	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.Update)
		apiV1.GET("/tags", tag.List)

		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.PATCH("/articles/:id/state", article.Update)
		apiV1.GET("/articles", article.List)
	}

	return r
}
