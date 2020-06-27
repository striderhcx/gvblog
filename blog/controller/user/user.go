package user

import (
	"blog/common"
	"blog/config"
	"blog/model"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var JsonResponse = common.JsonResponse

// Signin 用户登录
func Signin(c *gin.Context) {
	type LoginForm struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password" binding:"required,min=6,max=20"`
	}
	loginForm := LoginForm{}
	var nameOrEmail, password, sql string
	if err := c.ShouldBindWith(&loginForm, binding.JSON); err != nil {
		fmt.Println(err.Error())
		JsonResponse(common.ParamError, "参数有误", c)
		return
	}

	if loginForm.Email != "" {
		nameOrEmail = loginForm.Email
		password = loginForm.Password
		sql = "email = ?"
	} else if loginForm.Username != "" {
		nameOrEmail = loginForm.Username
		password = loginForm.Password
		sql = "username = ?"
	}

	user := model.User{}
	if err := model.DB.Where(sql, nameOrEmail).First(&user).Error; err != nil {
		JsonResponse(common.ParamError,"账号不存在", c)
		return
	}

	if user.CheckPassword(password) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id": user.ID,
			"exp": time.Now().Add(time.Second * time.Duration(config.TokenMaxAge)).Unix(),// 可以添加过期时间
		})
		tokenString, err := token.SignedString([]byte(config.TokenSecret))
		if err != nil {
			fmt.Println(err.Error())
			JsonResponse(common.SystemError,"内部错误",  c)
			return
		}
		//c.SetCookie("token", tokenString, config.TokenMaxAge, "/", "", false, true)

		c.JSON(http.StatusOK, gin.H{
			"code": common.Success,
			"msg":  "ok",
			"data": map[string]string{"token": tokenString},
		})
		return
	}
	JsonResponse(common.BusinessError,"账号或密码错误", c)
}

// Signup 用户注册
func Signup(c *gin.Context) {
	type UserForm struct {
		Username     string `form:"username" json:"username" binding:"required,min=4,max=20"`
		Email    string `form:"email" json:"email" binding:"email"`
		Password string `form:"password" json:"password" binding:"required,min=6,max=20"`
	}
	userFormData := UserForm{}
	if err := c.ShouldBindWith(&userFormData, binding.JSON); err != nil {
		fmt.Println(err.Error())
		JsonResponse(common.ParamError, fmt.Sprintf("参数无效: %s", err.Error()), c)
		return
	}
	userFormData.Username = strings.TrimSpace(userFormData.Username)
	userFormData.Email = strings.TrimSpace(userFormData.Email)
	if strings.Index(userFormData.Username, "@") != -1 {
		JsonResponse(common.ParamError, "用户名中不能含有@字符", c)
		return
	}
	user := model.User{}
	if err := model.DB.Where("email = ? OR username = ?", userFormData.Email, userFormData.Username).Find(&user).Error; err == nil {
		if user.Username == userFormData.Username {
			JsonResponse(common.ParamError, "用户名 "+user.Username+" 已被注册", c)
			return
		} else if user.Email == userFormData.Email {
			JsonResponse(common.ParamError, "邮箱 "+user.Email+" 已存在", c)
			return
		}
	}

	newUser := model.User{}
	newUser.Username = userFormData.Username
	newUser.Email = userFormData.Email
	newUser.Password = newUser.EncryptPassword(userFormData.Password, newUser.Salt())

	if err := model.DB.Create(&newUser).Error; err != nil {
		JsonResponse(common.MySqlError,"error", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": common.Success,
		"msg":  "ok",
		"data": nil,
	})
}

// Signout 退出登录
func Signout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": common.Success,
		"msg":  "ok",
		"data": nil,
	})
}
