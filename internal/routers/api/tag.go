package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mogfee/blog-server/global"
	"github.com/mogfee/blog-server/internal/model"
	"github.com/mogfee/blog-server/internal/service"
	"github.com/mogfee/blog-server/pkg/app"
	"github.com/mogfee/blog-server/pkg/errcode"
)

type TagController struct {
}

func NewTagController() TagController {
	return TagController{}
}
func (c TagController) Get(ctx *gin.Context) {

}

func (c TagController) List(ctx *gin.Context) {
}

func (c TagController) Create(ct *gin.Context) {
	ginWap := app.NewGinWap(ct)
	span, ctx := ginWap.StartSpanFromContext("添加标签控制器")
	defer span.Finish()
	ginWap.SetSpan(span)
	post := model.TagControllerCreateSwagger{}
	if err := ginWap.Bind(&post); err != nil {
		ginWap.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	//业务
	tagServer := service.NewTagServer(global.DBEngine)
	res, err := tagServer.Create(ctx, service.TagCreateParam{
		Name:  post.Name,
		State: post.State,
	})
	if err != nil {
		ginWap.ToErrorResponse(err)
		return
	}
	ginWap.ToResponse(gin.H{
		"id": res.Id,
	})
}
func (c TagController) Update(ctx *gin.Context) {

}

func (c TagController) Delete(ctx *gin.Context) {

}
