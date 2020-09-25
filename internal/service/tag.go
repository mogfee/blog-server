package service

import (
	"context"
	"github.com/mogfee/blog-server/global"
	"github.com/mogfee/blog-server/internal/dao"
	"github.com/mogfee/blog-server/internal/model"
	"github.com/mogfee/blog-server/pkg/errcode"
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
	global.Logger.WithCaller(1).WithContext(ctx).Infof("tagServer.Create: param:%+v", param)
	if param.Name == "" {
		global.Logger.WithCaller(1).WithContext(ctx).Error("tagServer.Create: error: 标签名不能为空")
		return nil, errcode.InvalidParams.WithDetails("标签名不能为空")
	}

	tagDao := dao.TagDao{}
	if t, err := tagDao.GetByName(ctx, s.db, param.Name); err != nil {
		return nil, err
	} else if t.Id > 0 {
		global.Logger.WithCaller(1).WithContext(ctx).Errorf("tagServer.Create: error: %s 标签名已存在", param.Name)
		return nil, errcode.InvalidParams.WithDetails("标签名已存在")
	}

	tag := model.Tag{
		Name:  param.Name,
		State: param.State,
	}
	err := tagDao.Create(ctx, s.db, &tag)
	if err != nil {
		global.Logger.WithCaller(1).WithContext(ctx).Errorf("tagServer.Create: tagDao.Create error: %s", err.Error())
		return nil, err
	}
	return &tag, err
}
