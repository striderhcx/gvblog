package model

import "time"

// Post 用户
type Post struct {
	ID           int       `gorm:"primary_key" form:"id" json:"id"`
	CategoryID   int       `form:"category_id" json:"category_id"`
	CreatedAt  	 time.Time  `json:"created_time"`
	UpdatedAt    time.Time  `json:"updated_time"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_time"`
	Title        string     `form:"title" json:"title"`
	Content      string     `form:"content" json:"content"`
	Pstatus      int        `json:"pstatus"`   // 是否删除，1删除，0未删除
}