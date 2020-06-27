package model

import (
	"blog/config"
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

// User 用户
type User struct {
	ID          int        `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `sql:"index" json:"deletedAt"`
	Username    string     `json:"username"`
	Password    string     `json:"-"`
	Email       string     `json:"-"`
	Pstatus     int        `json:"pstatus"`
}

// CheckPassword 验证密码是否正确
func (user User) CheckPassword(password string) bool {
	if password == "" || user.Password == "" {
		return false
	}
	return user.EncryptPassword(password, user.Salt()) == user.Password
}

// Salt 每个用户都有一个不同的盐
func (user User) Salt() string {
	var userSalt string
	if user.Password == "" {
		userSalt = strconv.Itoa(int(time.Now().Unix()))
	} else {
		userSalt = user.Password[0:10]
	}
	return userSalt
}

// EncryptPassword 给密码加密
func (user User) EncryptPassword(password, salt string) (hash string) {
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	hash = salt + password + config.PassSalt
	hash = salt + fmt.Sprintf("%x", md5.Sum([]byte(hash)))
	return
}

