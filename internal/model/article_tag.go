package model

type ArticleTag struct {


	Id         int64  `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	ModifiedOn int64  `json:"modified_on"`
	CreatedOn  int64  `json:"created_on"`
	DeletedOn  int64  `json:"deleted_on"`
	IsDel      int64  `json:"is_del"`


	TagId     int64 `json:"tag_id"`
	ArticleId int64 `json:"article_id"`
}

func (ArticleTag) TableName() string {
	return "blog_article_tag"
}
