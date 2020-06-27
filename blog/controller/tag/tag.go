package tag

import (
	"blog/common"
	"blog/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strings"
)

var JsonResponse = common.JsonResponse

func TagListView(c *gin.Context) {
	tags := []model.Tag{}
	model.DB.Order("created_at desc, id desc").Find(&tags)

	tagMap := map[int]interface{}{}
	for _, tag := range tags {
		tagMap[tag.ID] = tag.Name
	}
	type Result struct {
		TagId  int
		Total int
	}
	var results []Result
	// .Having("total > 0")
	model.DB.Model(model.PostTag{}).Select("tag_id, count(*) as  total").Group("tag_id").Scan(&results)

	tagPostCntMap := map[int]interface{}{}
	for _, item := range results {
		tagPostCntMap[item.TagId] = item.Total
	}

	respTags := []map[string]interface{}{}
	for _, tag := range tags {
		postCount, _ := tagPostCntMap[tag.ID]
		tagItem := map[string]interface{}{
			"id":            tag.ID,
			"name":          tag.Name,
			"post_count":	 postCount,
		}
		respTags = append(respTags, tagItem)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":       common.Success,
		"msg":        "ok",
		"data":       respTags,
		"pagination": nil,
	})
}

func Create(c *gin.Context) {
	// todo 自定义错误提示返回
	type TagParams struct {
		Name        string     `form:"name" binding:"required,min=1,max=40"`
	}
	tagForm := TagParams{}
	if err := c.ShouldBindWith(&tagForm, binding.Form); err != nil {
		fmt.Println(err.Error())
		JsonResponse(common.ParamError,"参数无效", c)
		return
	}

	newTag := model.Tag{}
	newTag.Name = strings.TrimSpace(tagForm.Name)

	saveErr := model.DB.Create(&newTag).Error
	if saveErr != nil {
		fmt.Println(saveErr.Error())
		JsonResponse(common.MySqlError,"error", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": common.Success,
		"msg":  "ok",
		"data": nil,
	})
}
