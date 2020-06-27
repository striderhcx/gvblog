package model

import "time"

// post Tag
type Tag struct {
	ID           int       `gorm:"primary_key" form:"id" json:"id"`
	Name         string     `form:"name" json:"name"`

	CreatedAt  	 time.Time  `json:"created_time"`
	UpdatedAt    time.Time  `json:"updated_time"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_time"`
	Pstatus      int        `json:"pstatus"`   // 是否删除，1删除，0未删除
}
