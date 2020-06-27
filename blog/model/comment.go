package model

import (
	"time"
)

// Comment 评论
type Comment struct {
	ID           int       `gorm:"primary_key" json:"id"`
	PostID	 	 uint       `form:"post_id" json:"post_id"`
	Username     string 	`form:"username" json:"username"`
	Content      string     `form:"content" json:"content"`
	Pstatus      int        `json:"pstatus"`   // 是否删除，1删除，0未删除
	CreatedAt  	 time.Time  `json:"created_time"`
	UpdatedAt    time.Time  `json:"updated_time"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_time"`
}