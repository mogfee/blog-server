package dao

import (
	"context"
	"errors"
	"github.com/mogfee/blog-server/internal/model"
	"gorm.io/gorm"
)

type TagDao struct {
	Id int64
}

func (m TagDao) Create(ctx context.Context, tx *gorm.DB, data *model.Tag) error {
	return tx.WithContext(ctx).Create(&data).Error
}

func (m TagDao) GetByName(ctx context.Context, tx *gorm.DB, tagName string) (model.Tag, error) {
	if tagName == "" {
		return model.Tag{}, errors.New("标签名称不能为空")
	}
	row := model.Tag{}
	err := tx.WithContext(ctx).Where("name=? and is_del=0", tagName).First(&row).Error
	if err == gorm.ErrRecordNotFound {
		return row, nil
	}
	return row, err
}

func (m TagDao) Update(ctx context.Context, tx *gorm.DB, updates map[string]interface{}) error {
	if m.Id <= 0 {
		return errors.New("标签ID不能为空")
	}
	return tx.WithContext(ctx).Table(model.Tag{}.TableName()).Where("id=?", m.Id).Updates(updates).Error
}
