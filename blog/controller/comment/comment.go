package comment

import (
	"blog/common"
	"blog/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var JsonResponse = common.JsonResponse

func Create(c *gin.Context) {

	type CommentParams struct {
		PostID   int       `json:"post_id" binding:"required,numeric,min=1"`
		UserName        string     `json:"username" binding:"required,min=2,max=40"`
		Content      string     `json:"content" binding:"required"`
	}
	commentForm := CommentParams{}
	if err := c.ShouldBindWith(&commentForm, binding.JSON); err != nil {
		fmt.Println(err.Error())
		JsonResponse(common.ParamError, "参数无效", c)
		return
	}
	commentForm.Content = strings.TrimSpace(commentForm.Content)
	if commentForm.Content == "" {
		JsonResponse(common.ParamError, "评论内容不能为空", c)
		return
	}

	createErr := model.DB.Create(&commentForm).Error
	if createErr != nil {
		fmt.Println(createErr.Error())
		JsonResponse(common.MySqlError, "error", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": common.Success,
		"msg":  "ok",
		"data": nil,
	})
}

func Update(c *gin.Context) {
	type CommentUpdateForm struct {
		CommentId  int      `json:"comment_id" binding:"required"`
		Content    string   `json:"content" binding:"required"`
	}
	commentUpdateForm := CommentUpdateForm{}
	if err := c.ShouldBindWith(&commentUpdateForm, binding.JSON); err != nil {
		fmt.Println(err.Error())
		JsonResponse(common.ParamError,"参数无效", c)
		return
	}
	comment := model.Comment{}
	if model.DB.First(&comment, commentUpdateForm.CommentId).Error != nil {
		JsonResponse(common.ParamError, fmt.Sprintf("不存在ID为%v的评论", commentUpdateForm.CommentId), c)
		return
	}
	commentUpdateForm.Content = strings.TrimSpace(commentUpdateForm.Content)
	if commentUpdateForm.Content == "" {
		JsonResponse(common.ParamError, "评论不能为空", c)
		return
	}
	updateErr := model.DB.Model(&comment).Updates(map[string]interface{}{"content": commentUpdateForm.Content}).Error
	if updateErr != nil {
		fmt.Println(updateErr.Error())
		JsonResponse(common.MySqlError,"error", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": common.Success,
		"msg":  "ok",
		"data": nil,
	})
}

func Delete(c *gin.Context) {
	type CommentDeleteForm struct {
		Id  int      `json:"id" binding:"required"`
	}
	commentDeleteForm := CommentDeleteForm{}
	if err := c.ShouldBindWith(&commentDeleteForm, binding.JSON); err != nil {
		fmt.Println(err.Error())
		JsonResponse(common.ParamError,"参数无效", c)
		return
	}
	comment := model.Comment{}
	if model.DB.First(&comment, commentDeleteForm.Id).Error != nil {
		JsonResponse(common.ParamError,fmt.Sprintf("不存在ID为%v的评论", commentDeleteForm.Id), c)
		return
	}
	deleteErr := model.DB.Delete(&comment).Error
	if deleteErr != nil {
		fmt.Println(deleteErr.Error())
		JsonResponse(common.MySqlError,"error", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": common.Success,
		"msg":  "ok",
		"data": nil,
	})
}

func GetCommentListView(c *gin.Context) {
	type CommentListParam struct {
		PostId      int    `form:"post_id" binding:"required"`
		Page		int    `form:"page"`
		Limit       int    `form:"limit" binding:"required"`
	}
	params := CommentListParam{}
	if err := c.ShouldBindQuery(&params); err != nil {
		JsonResponse(common.ParamError,"请求参数有误", c)
		fmt.Println(err)
		return
	}
	post_id := params.PostId
	pageComments := []model.Comment{}

	model.DB.Where("post_id = ?", post_id).
		Order("created_at asc, id asc").
		Offset(params.Page * params.Limit).
		Limit(params.Limit).
		Find(&pageComments)

	count := 0
	model.DB.Table("comments").Where("pstatus= ? ", model.DataStatusNormal).Count(&count)
	pagination := common.Pagination(count, params.Page, params.Limit)
	commentSlice := []map[string]interface{}{}
	for _, comment := range pageComments {
		item := map[string]interface{}{
			"id":           comment.ID,
			"content":      comment.Content,
			"username":  comment.Username,
			"created_time": comment.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		commentSlice = append(commentSlice, item)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":       common.Success,
		"msg":        "ok",
		"data":       commentSlice,
		"pagination": pagination,
	})
}
