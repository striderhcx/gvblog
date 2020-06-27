package model

import (
	"blog/config"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"time"
)

// DB 数据库连接
var DB *gorm.DB

// RedisPool Redis链接池
var RedisPool *redis.Pool


func initDB() {
	db, err := gorm.Open(config.Dialect, fmt.Sprintf("%s:%s@tcp(%s:%d)/blog?charset=utf8mb4&parseTime=true&loc=Asia%sShanghai", config.DBConfig.User, config.DBConfig.Password, config.DBConfig.Host, config.DBConfig.Port, `%2f`))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	DB = db
}

func initRedis() {
	config.RedisConfig.URL = fmt.Sprintf("%s:%d", config.RedisConfig.Host, config.RedisConfig.Port)
	RedisPool = &redis.Pool{
		MaxIdle:     config.RedisConfig.MaxIdle,
		MaxActive:   config.RedisConfig.MaxActive,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.RedisConfig.URL, redis.DialPassword(config.RedisConfig.Password))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

func init() {
	initDB()
	initRedis()
}