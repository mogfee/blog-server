package model

type Article struct {


	Id         int64  `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	ModifiedOn int64  `json:"modified_on"`
	CreatedOn  int64  `json:"created_on"`
	DeletedOn  int64  `json:"deleted_on"`
	IsDel      int64  `json:"is_del"`


	State         int64  `json:"state"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
}

func (Article) TableName() string {
	return "blog_article"
}
