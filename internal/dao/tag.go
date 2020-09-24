package dao

import (
	"context"
	"github.com/mogfee/blog-server/internal/model"
	"gorm.io/gorm"
)

type TagDao struct {
	*model.Tag
}

func (m TagDao) Create(ctx context.Context, tx *gorm.DB, data *model.Tag) (*model.Tag, error) {
	return data.Create(ctx, tx)
}
