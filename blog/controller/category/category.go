package category

import (
	utils "blog/common"
	"blog/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strings"
)

var JsonResponse = utils.JsonResponse

func CategoryListView(c *gin.Context) {
	categories := []model.Category{}
	model.DB.Find(&categories)

	type Result struct {
		CategoryId  int
		Total int
	}
	var results []Result
	model.DB.Model(model.Post{}).Select("category_id, count(*) as  total").Group("category_id").Scan(&results)

	categoryCountMap := map[int]int{}
	for _, item := range results {
		categoryCountMap[item.CategoryId] = item.Total
	}

	respCategories := []map[string]interface{}{}
	for _, cate := range categories {
		categoryItem := map[string]interface{}{
			"id":           cate.ID,
			"name":         cate.Name,
			"post_count":   categoryCountMap[cate.ID],
			"created_time": cate.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		respCategories = append(respCategories, categoryItem)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": utils.Success,
		"msg":  "ok",
		"data": respCategories,
	})
}

func Create(c *gin.Context) {
	type CategoryCreateForm struct {
		Name        string   `form:"name" binding:"required"`
	}
	categoryCreateForm := CategoryCreateForm{}
	if err := c.ShouldBindWith(&categoryCreateForm, binding.Form); err != nil {
		fmt.Println(err.Error())
		JsonResponse(utils.ParamError, "参数无效", c)
		return
	}
	categoryCreateForm.Name = strings.TrimSpace(categoryCreateForm.Name)
	if categoryCreateForm.Name == "" {
		JsonResponse(utils.ParamError, "类型名不能为空", c)
		return
	}
	newCategory := model.Category{}
	newCategory.Name = categoryCreateForm.Name
	createErr := model.DB.Create(&newCategory).Error
	if createErr != nil {
		fmt.Println(createErr.Error())
		JsonResponse(utils.MySqlError, "error", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": utils.Success,
		"msg":  "ok",
		"data": nil,
	})
}

func Update(c *gin.Context) {
	type CategoryUpdateForm struct {
		CategoryId      int       `form:"category_id" binding:"required"`
		Name             string    `form:"name" binding:"required"`
	}
	categoryUpdateForm := CategoryUpdateForm{}
	if err := c.ShouldBindWith(&categoryUpdateForm, binding.Form); err != nil {
		fmt.Println(err.Error())
		JsonResponse(utils.ParamError, "参数无效", c)
		return
	}
	category := model.Category{}
	if model.DB.First(&category, categoryUpdateForm.CategoryId).Error != nil {
		JsonResponse(utils.ParamError, fmt.Sprintf("不存在ID为%v的博文", categoryUpdateForm.CategoryId), c)
		return
	}

	if categoryUpdateForm.Name == "" {
		JsonResponse(utils.ParamError,"类型名不能为空",  c)
		return
	}
	categoryUpdateForm.Name = strings.TrimSpace(categoryUpdateForm.Name)
	updateErr := model.DB.Model(&category).Updates(map[string]interface{}{"name": categoryUpdateForm.Name}).Error
	if updateErr != nil {
		fmt.Println(updateErr.Error())
		JsonResponse(utils.MySqlError,"error", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": utils.Success,
		"msg":  "ok",
		"data": nil,
	})
}

func Delete(c *gin.Context) {
	type CategoryDeleteForm struct {
		CategoryId     int    `form:"category_id" binding:"required"`
	}
	categoryDeleteForm := CategoryDeleteForm{}
	if err := c.ShouldBindWith(&categoryDeleteForm, binding.Form); err != nil {
		fmt.Println(err.Error())
		JsonResponse(utils.ParamError, "参数无效", c)
		return
	}
	category := model.Category{}
	if model.DB.First(&category, categoryDeleteForm.CategoryId).Error != nil {
		JsonResponse(utils.ParamError, fmt.Sprintf("不存在ID为%v的分类", categoryDeleteForm.CategoryId), c)
		return
	}
	deleteErr := model.DB.Delete(&category).Error
	if deleteErr != nil {
		fmt.Println(deleteErr.Error())
		JsonResponse(utils.MySqlError, "error", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": utils.Success,
		"msg":  "ok",
		"data": nil,
	})
}
