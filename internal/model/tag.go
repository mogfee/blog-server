package model

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

type TagControllerCreateSwagger struct {
	Name  string `json:"name" binding:"required"`
	State int64  `json:"state"`
}
