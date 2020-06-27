package model

import (
	"time"
)

// Category 分类
type Category struct {
	ID           int       `gorm:"primary_key" json:"id"`
	Name         string     `form:"name" json:"name"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `sql:"index" json:"deletedAt"`
}