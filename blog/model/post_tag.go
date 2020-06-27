package model

import "time"

type PostTag struct {
	ID           int       `gorm:"primary_key" form:"id" json:"id"`
	PostID   	 int       `form:"post_id" json:"post_id"`
	TagId		 int       `form:"tag_id" json:"tag_id"`
	CreatedAt  	 time.Time  `json:"created_time"`
	UpdatedAt    time.Time  `json:"updated_time"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_time"`
	Pstatus      int        `json:"pstatus"`   // 是否删除，1删除，0未删除
}
