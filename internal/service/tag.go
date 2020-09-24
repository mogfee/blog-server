package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/mogfee/blog-server/internal/dao"
	"github.com/mogfee/blog-server/internal/model"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

type tagServer struct {
	db *gorm.DB
}

func NewTagServer(db *gorm.DB) *tagServer {
	return &tagServer{db: db}
}

type TagCreateParam struct {
	Name  string
	State int64
}

func (s tagServer) Create(ctx context.Context, param TagCreateParam) (*model.Tag, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "tagServer.Create")
	defer span.Finish()
	span.SetBaggageItem("in", fmt.Sprintf("%+v", param))
	//global.Logger.WithCaller(1).WithTrace(span).Infof( "param:%+v", param)
	if param.Name == "" {
		return nil, errors.New("标签名不能为空")
	}

	tag := model.Tag{
		Name:  param.Name,
		State: param.State,
	}
	return dao.TagDao{}.Create(ctx, s.db, &tag)
}
