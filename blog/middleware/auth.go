package middleware

import (
	"blog/common"
	"blog/config"
	"blog/model"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JsonResponse = common.JsonResponse


func getUser(c *gin.Context) (model.User, error) {
	var user model.User
	tokenString := c.GetHeader("Authorization")
	fmt.Println("传过来的token", tokenString)
	token, tokenErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.TokenSecret), nil
	})

	if tokenErr != nil {
		return user, errors.New("未登录")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["id"].(float64))
		user.ID = userID
		return user, nil
	}
	return user, errors.New("未登录")
}

// SetContextUser 给 context 设置 user
func SetContextUser(c *gin.Context) {
	var user model.User
	var err error
	if user, err = getUser(c); err != nil {
		c.Set("user", nil)
		c.Next()
		return
	}
	c.Set("user", user)
	c.Next()
}

// LoginRequired 必须是登录用户
func LoginRequired(c *gin.Context) {
	var user model.User
	var err error
	if user, err = getUser(c); err != nil {
		JsonResponse(common.AuthError, "未登录", c)
		return
	}
	c.Set("user", user)
	c.Next()
}
