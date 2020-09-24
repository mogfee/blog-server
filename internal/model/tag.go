package model

import (
	"context"
	"errors"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

type Tag struct {
	Id         int64  `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	ModifiedOn int64  `json:"modified_on"`
	CreatedOn  int64  `json:"created_on"`
	DeletedOn  int64  `json:"deleted_on"`
	IsDel      int64  `json:"is_del"`

	Name  string `json:"name"`
	State int64  `json:"state"`
}

func (Tag) TableName() string {
	return "blog_tag"
}
func (m Tag) Create(ctx context.Context, tx *gorm.DB) (*Tag, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Tag.Create")
	defer span.Finish()
	err := tx.WithContext(ctx).Create(&m).Error
	return &m, err
}

func (m Tag) GetByName(ctx context.Context, tx *gorm.DB) (Tag, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Tag.GetByName")
	defer span.Finish()

	if m.Name == "" {
		span.SetBaggageItem("error", "标签名称不能为空")
		return Tag{}, errors.New("标签名称不能为空")
	}

	row := Tag{}
	err := tx.WithContext(ctx).Where("name=? and is_del=0", m.Name).First(&row).Error
	return row, err
}

func (m Tag) Update(ctx context.Context, tx *gorm.DB, updates map[string]interface{}) (Tag, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Tag.Update")
	defer span.Finish()
	row := Tag{}
	err := tx.WithContext(ctx).Where("name=? and is_del=0", m.Name).First(&row).Error
	return row, err
}
