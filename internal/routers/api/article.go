package api

import "github.com/gin-gonic/gin"

type Article struct {
}

func NewArticle() Article {
	return Article{}
}
func (c Article) Get(ctx *gin.Context) {

}

func (c Article) List(ctx *gin.Context) {

}
func (c Article) Create(ctx *gin.Context) {

}
func (c Article) Update(ctx *gin.Context) {

}

func (c Article) Delete(ctx *gin.Context) {

}
