package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/mogfee/blog-server/pkg/errcode"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"net/http"
)

type response struct {
	Ctx     *gin.Context
	spanCtx opentracing.Span
}
type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewGinWap(ctx *gin.Context) *response {
	return &response{
		Ctx: ctx,
	}
}
func (r *response) SetSpan(span opentracing.Span) {
	r.spanCtx = span
}
func (r *response) Bind(data interface{}) error {
	//span, _ := r.StartSpanFromContext("bind")
	//span.SetTag("error", false)
	b := binding.Default(r.Ctx.Request.Method, r.Ctx.ContentType())
	if err := r.Ctx.ShouldBindWith(data, b); err != nil {
		return err
	}
	// 参数验证
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return err
	}
	r.spanCtx.SetBaggageItem("bind_in", fmt.Sprintf("%+v", data))
	return nil
}
func (r *response) StartSpanFromContext(operationName string) (opentracing.Span, context.Context) {

	spanContext, ok := r.Ctx.Get("tracer_span")
	if ok {
		sp := spanContext.(*jaeger.Span)
		span, ctx := opentracing.StartSpanFromContext(r.Ctx, operationName, opentracing.ChildOf(sp.Context()))
		return span, ctx
	}
	return opentracing.StartSpanFromContext(r.Ctx, operationName)
}

func (r *response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	resp := gin.H{"code": 0, "msg": "ok", "data": data}
	r.spanCtx.SetTag("error", false)
	r.spanCtx.SetBaggageItem("resp", fmt.Sprintf("%+v", resp))
	r.Ctx.JSON(http.StatusOK, resp)
}
func (r *response) ToResponseList(list interface{}, totalRows int) {
	resp := gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	}
	r.spanCtx.SetTag("error", false)
	r.spanCtx.SetBaggageItem("resp", fmt.Sprintf("%+v", resp))
	r.Ctx.JSON(http.StatusOK, resp)
}

func (r *response) ToErrorResponse(e error) {
	var err = &errcode.Error{}
	switch e.(type) {
	case *errcode.Error:
		err = e.(*errcode.Error)
	default:
		err = errcode.ServerError.WithDetails(e.Error())
	}
	resp := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		resp["details"] = details
	}
	r.spanCtx.SetTag("error", true)
	r.spanCtx.SetBaggageItem("resp", fmt.Sprintf("%+v", resp))
	r.Ctx.JSON(err.StatusCode(), resp)
}

//func (r *response) ToErrorResponse(err *errcode.Error) {
//	resp := gin.H{"code": err.Code(), "msg": err.Msg()}
//	details := err.Details()
//	if len(details) > 0 {
//		resp["details"] = details
//	}
//	r.spanCtx.SetTag("error", true)
//	r.spanCtx.SetBaggageItem("resp", fmt.Sprintf("%+v", resp))
//	r.Ctx.JSON(err.StatusCode(), resp)
//}
